package connectfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/handler/handlerfx"
	sconnect "github.com/ScoreTrak/ScoreTrak/pkg/server/connect"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlerfx.ConnectModule,
	fx.Provide(
		fx.Annotate(sconnect.NewConnectServer, fx.ParamTags(`group:connect-handlers`)),
	),
	fx.Invoke(sconnect.InitConnectServer),
)
