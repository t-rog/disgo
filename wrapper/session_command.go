// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.

package wrapper

// Command sends an Opcode 1 Heartbeat command to the Discord Gateway.
func (c *Heartbeat) Command(session *Session) error {
	if err := writeEvent(session, FlagGatewayOpcodeHeartbeat, FlagGatewayCommandNameHeartbeat, c); err != nil {
		return err
	}

	return nil
}

// Command sends an Opcode 2 Identify command to the Discord Gateway.
func (c *Identify) Command(session *Session) error {
	if err := writeEvent(session, FlagGatewayOpcodeIdentify, FlagGatewayCommandNameIdentify, c); err != nil {
		return err
	}

	return nil
}

// Command sends an Opcode 3 PresenceUpdate command to the Discord Gateway.
func (c *GatewayPresenceUpdate) Command(session *Session) error {
	if err := writeEvent(session, FlagGatewayOpcodePresenceUpdate, FlagGatewayCommandNamePresenceUpdate, c); err != nil {
		return err
	}

	return nil
}

// Command sends an Opcode 4 VoiceStateUpdate command to the Discord Gateway.
func (c *VoiceStateUpdate) Command(session *Session) error {
	if err := writeEvent(session, FlagGatewayOpcodeVoiceStateUpdate, FlagGatewayCommandNameVoiceStateUpdate, c); err != nil {
		return err
	}

	return nil
}

// Command sends an Opcode 6 Resume command to the Discord Gateway.
func (c *Resume) Command(session *Session) error {
	if err := writeEvent(session, FlagGatewayOpcodeResume, FlagGatewayCommandNameResume, c); err != nil {
		return err
	}

	return nil
}

// Command sends an Opcode 8 RequestGuildMembers command to the Discord Gateway.
func (c *RequestGuildMembers) Command(session *Session) error {
	if err := writeEvent(session, FlagGatewayOpcodeRequestGuildMembers, FlagGatewayCommandNameRequestGuildMembers, c); err != nil {
		return err
	}

	return nil
}
