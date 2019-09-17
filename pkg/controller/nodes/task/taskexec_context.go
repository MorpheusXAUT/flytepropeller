package task

import (
	"bytes"
	"context"
	"strconv"

	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

	pluginCatalog "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/catalog"
	pluginCore "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/io"

	"github.com/lyft/flytepropeller/pkg/controller/nodes/errors"
	"github.com/lyft/flytepropeller/pkg/controller/nodes/handler"
	"github.com/lyft/flytepropeller/pkg/controller/nodes/task/catalog"
	"github.com/lyft/flytepropeller/pkg/utils"
)

const IDMaxLength = 50

type taskExecutionID struct {
	execName string
	id       *core.TaskExecutionIdentifier
}

func (te taskExecutionID) GetID() core.TaskExecutionIdentifier {
	return *te.id
}

func (te taskExecutionID) GetGeneratedName() string {
	return te.execName
}

type taskExecutionMetadata struct {
	handler.NodeExecutionMetadata
	taskExecID taskExecutionID
	o          pluginCore.TaskOverrides
}

func (t taskExecutionMetadata) GetTaskExecutionID() pluginCore.TaskExecutionID {
	return t.taskExecID
}

func (t taskExecutionMetadata) GetOverrides() pluginCore.TaskOverrides {
	return t.o
}

type taskExecutionContext struct {
	handler.NodeExecutionContext
	tm  taskExecutionMetadata
	rm  pluginCore.ResourceManager
	psm *pluginStateManager
	tr  handler.TaskReader
	ow  *OutputWriter
	ber *bufferedEventRecorder
	sm  pluginCore.SecretManager
	c   catalog.Client
}

func (t taskExecutionContext) EventsRecorder() pluginCore.EventsRecorder {
	return t.ber
}

func (t taskExecutionContext) ResourceManager() pluginCore.ResourceManager {
	return t.rm
}

func (t taskExecutionContext) PluginStateReader() pluginCore.PluginStateReader {
	return t.psm
}

func (t *taskExecutionContext) TaskReader() pluginCore.TaskReader {
	return t.tr
}

func (t *taskExecutionContext) TaskExecutionMetadata() pluginCore.TaskExecutionMetadata {
	return t.tm
}

func (t *taskExecutionContext) OutputWriter() io.OutputWriter {
	return t.ow
}

func (t *taskExecutionContext) PluginStateWriter() pluginCore.PluginStateWriter {
	return t.psm
}

func (t *taskExecutionContext) Catalog() pluginCatalog.Client {
	return nil
}

func (t taskExecutionContext) SecretManager() pluginCore.SecretManager {
	return t.sm
}

func (t *Handler) newTaskExecutionContext(ctx context.Context, nCtx handler.NodeExecutionContext) (*taskExecutionContext, error) {

	id := GetTaskExecutionIdentifier(nCtx)

	uniqueID, err := utils.FixedLengthUniqueIDForParts(IDMaxLength, nCtx.NodeExecutionMetadata().GetOwnerID().Name, nCtx.NodeID(), strconv.Itoa(int(id.RetryAttempt)))
	if err != nil {
		// SHOULD never really happen
		return nil, err
	}

	ow, err := NewRemoteFileOutputWriter(ctx, nCtx.NodeStatus().GetDataDir(), nCtx.DataStore())
	if err != nil {
		return nil, err
	}

	ts := nCtx.NodeStateReader().GetTaskNodeState()
	var b *bytes.Buffer
	if ts.PluginState != nil {
		b = bytes.NewBuffer(ts.PluginState)
	}
	psm, err := newPluginStateManager(ctx, GobCodecVersion, ts.PluginStateVersion, b)
	if err != nil {
		return nil, errors.Wrapf(errors.RuntimeExecutionError, nCtx.NodeID(), err, "unable to initialize plugin state manager")
	}

	return &taskExecutionContext{
		NodeExecutionContext: nCtx,
		tm: taskExecutionMetadata{
			NodeExecutionMetadata: nCtx.NodeExecutionMetadata(),
			taskExecID:            taskExecutionID{execName: uniqueID, id: id},
			o:                     nCtx.Node(),
		},
		// TODO add resource manager
		rm:  nil,
		psm: psm,
		tr:  nCtx.TaskReader(),
		ow:  ow,
		ber: newBufferedEventRecorder(),
		c:   t.catalog,
		// TODO @kumare path should be configurable
		sm: fileEnvSecretManager{"/etc/secrets"},
	}, nil
}