package etherstate

//// NewHandler ...
//func NewHandler(k keeper.Keeper) sdk.Handler {
//	_ = keeper.NewMsgServerImpl(k)
//
//	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
//		ctx = ctx.WithEventManager(sdk.NewEventManager())
//
//		switch msg := msg.(type) {
//		case *types.MsgDisableEthAddress:
//		case *types.MsgEnableEthAddress:
//			// do nothing
//			return
//		default:
//			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
//			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
//		}
//	}
//}
