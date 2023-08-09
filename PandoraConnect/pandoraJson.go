package PandoraConnect

type PandoraAuthorizeJson struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	UserID    int    `json:"user_id"`
	Lang      string `json:"lang"`
	SessionID string `json:"session_id"`
}
type PandoraPingJson struct {
	Status string `json:"status"`
}
type PandoraDevicesList []struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	CarType   int    `json:"car_type"`
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	Color     string `json:"color"`
	AutoMarka string `json:"auto_marka"`
	AutoModel string `json:"auto_model"`
	Tanks     []any  `json:"tanks"`
	Features  struct {
		Heater         int `json:"heater"`
		AutoCheck      int `json:"auto_check"`
		Beep           int `json:"beep"`
		Events         int `json:"events"`
		Value100       int `json:"value_100"`
		Channel        int `json:"channel"`
		Sensors        int `json:"sensors"`
		Tracking       int `json:"tracking"`
		Connection     int `json:"connection"`
		ActiveSecurity int `json:"active_security"`
		Notification   int `json:"notification"`
		Schedule       int `json:"schedule"`
		Autostart      int `json:"autostart"`
		Light          int `json:"light"`
		Trunk          int `json:"trunk"`
		KeepAlive      int `json:"keep_alive"`
		CustomPhones   int `json:"custom_phones"`
		ExtendProps    int `json:"extend_props"`
		Bluetooth      int `json:"bluetooth"`
		ObdCodes       int `json:"obd_codes"`
		HeaterFrom40   int `json:"heater_from_40"`
	} `json:"features"`
	FuelTank    int `json:"fuel_tank"`
	Permissions struct {
		Control      int `json:"control"`
		Settings     int `json:"settings"`
		SettingsSave int `json:"settings_save"`
		Events       int `json:"events"`
		Tracks       int `json:"tracks"`
		Status       int `json:"status"`
		Oauth        int `json:"oauth"`
		Rules        int `json:"rules"`
		Tanks        int `json:"tanks"`
		TanksSave    int `json:"tanks_save"`
		Detach       int `json:"detach"`
		Alarms       int `json:"alarms"`
	} `json:"permissions"`
	IsShared       bool   `json:"is_shared"`
	Phone          string `json:"phone"`
	Phone1         string `json:"phone1"`
	ActiveSim      int    `json:"active_sim"`
	Model          string `json:"model"`
	VoiceVersion   string `json:"voice_version"`
	Firmware       string `json:"firmware"`
	StartOwnership int    `json:"start_ownership"`
	OwnerID        int    `json:"owner_id"`
}

type PandoraUserProfileJson struct {
	Response struct {
		ID       int    `json:"id"`
		NameF    string `json:"name_f"`
		NameI    string `json:"name_i"`
		NameO    string `json:"name_o"`
		Balance  int    `json:"balance"`
		Demo     int    `json:"demo"`
		Lang     string `json:"lang"`
		Type     string `json:"type"`
		Features struct {
			Alarms struct {
				Type string `json:"type"`
				Val  int    `json:"val"`
			} `json:"alarms"`
			Trackers struct {
				Type string `json:"type"`
				Val  int    `json:"val"`
			} `json:"trackers"`
			Gz struct {
				Type string `json:"type"`
				Val  int    `json:"val"`
			} `json:"gz"`
			Rules struct {
				Type string `json:"type"`
				Val  int    `json:"val"`
			} `json:"rules"`
			Reports struct {
				Type string `json:"type"`
				Val  int    `json:"val"`
			} `json:"reports"`
			Schedule struct {
				Type string `json:"type"`
				Val  bool   `json:"val"`
			} `json:"schedule"`
			Sms struct {
				Type string `json:"type"`
				Val  int    `json:"val"`
			} `json:"sms"`
		} `json:"features"`
		CompanyName string `json:"company_name"`
		TempUnit    string `json:"temp_unit"`
		LenUnit     string `json:"len_unit"`
		CompanyID   int    `json:"company_id"`
	} `json:"response"`
}

// PandoraUpdatesJson 1079312388 следует заменить на ваш id сигнализации, при первом запуске программа его покажет, ниже в структуре он встречается несколько раз, впрочем, и без него будет работать.
type PandoraUpdatesJson struct {
	Ts    int   `json:"ts"`
	Lenta []any `json:"lenta"`
	Time  struct {
		Num1079312388 struct {
			Online  int `json:"online"`
			Onlined int `json:"onlined"`
			Command int `json:"command"`
			Setting int `json:"setting"`
		} `json:"1079312388"`
	} `json:"time"`
	Stats struct {
		Num1079312388 struct {
			Online       int     `json:"online"`
			Move         int     `json:"move"`
			Dtime        int     `json:"dtime"`
			DtimeRec     int     `json:"dtime_rec"`
			Voltage      float64 `json:"voltage"`
			EngineTemp   int     `json:"engine_temp"`
			LiquidSensor []any   `json:"liquid_sensor"`
			X            float64 `json:"x"`
			Y            float64 `json:"y"`
			BitState1    int     `json:"bit_state_1"`
			OutTemp      int     `json:"out_temp"`
			Balance      struct {
				Value int    `json:"value"`
				Cur   string `json:"cur"`
			} `json:"balance"`
			Balance1 struct {
				Value int    `json:"value"`
				Cur   string `json:"cur"`
			} `json:"balance1"`
			Sims []struct {
				PhoneNumber string `json:"phoneNumber"`
				IsActive    bool   `json:"isActive"`
				Balance     struct {
					Value int    `json:"value"`
					Cur   string `json:"cur"`
				} `json:"balance"`
			} `json:"sims"`
			ActiveSim  int     `json:"active_sim"`
			Speed      int     `json:"speed"`
			Tanks      []any   `json:"tanks"`
			EngineRpm  int     `json:"engine_rpm"`
			Rot        int     `json:"rot"`
			Fuel       int     `json:"fuel"`
			CabinTemp  int     `json:"cabin_temp"`
			Evaq       int     `json:"evaq"`
			GsmLevel   int     `json:"gsm_level"`
			Props      []any   `json:"props"`
			Mileage    float64 `json:"mileage"`
			MileageCAN int     `json:"mileage_CAN"`
			Metka      int     `json:"metka"`
			Brelok     int     `json:"brelok"`
			Relay      int     `json:"relay"`
			Smeter     int     `json:"smeter"`
			Tconsum    int     `json:"tconsum"`
			BenishMode bool    `json:"benish_mode"`
			Land       int     `json:"land"`
			Bunker     int     `json:"bunker"`
			ExStatus   int     `json:"ex_status"`
			Can        struct {
				CANMileageToMaintenance int  `json:"CAN_mileage_to_maintenance"`
				CANMileageToEmpty       int  `json:"CAN_mileage_to_empty"`
				CANDaysToMaintenance    int  `json:"CAN_days_to_maintenance"`
				CANConsumption          int  `json:"CAN_consumption"`
				CANConsumptionAfter     int  `json:"CAN_consumption_after"`
				CANAverageSpeed         int  `json:"CAN_average_speed"`
				CANTMPSReserve          int  `json:"CAN_TMPS_reserve"`
				CANTMPSBackLeft         int  `json:"CAN_TMPS_back_left"`
				CANTMPSBackRight        int  `json:"CAN_TMPS_back_right"`
				CANTMPSForvardLeft      int  `json:"CAN_TMPS_forvard_left"`
				CANTMPSForvardRight     int  `json:"CAN_TMPS_forvard_right"`
				CANLowLiquid            bool `json:"CAN_low_liquid"`
				CANDriverGlass          bool `json:"CAN_driver_glass"`
				CANBackLeftGlass        bool `json:"CAN_back_left_glass"`
				CANPassengerGlass       bool `json:"CAN_passenger_glass"`
				CANBackRightGlass       bool `json:"CAN_back_right_glass"`
				CANDriverBelt           bool `json:"CAN_driver_belt"`
				CANBackLeftBelt         bool `json:"CAN_back_left_belt"`
				CANPassengerBelt        bool `json:"CAN_passenger_belt"`
				CANBackRightBelt        bool `json:"CAN_back_right_belt"`
				CANBackCenterBelt       bool `json:"CAN_back_center_belt"`
				CANSeatTaken            bool `json:"CAN_seat_taken"`
				CANNeedPadsExchange     bool `json:"CAN_need_pads_exchange"`
				CANMileageByBattery     int  `json:"CAN_mileage_by_battery"`
				ChargingConnect         bool `json:"charging_connect"`
				ChargingSlow            bool `json:"charging_slow"`
				ChargingFast            bool `json:"charging_fast"`
				EvStatusReady           bool `json:"ev_status_ready"`
				Soh                     int  `json:"SOH"`
				Soc                     int  `json:"SOC"`
				BatteryTemperature      int  `json:"battery_temperature"`
			} `json:"can"`
			Heater struct {
				HeaterPower       int   `json:"heater_power"`
				HeaterVoltage     int   `json:"heater_voltage"`
				HeaterFlame       bool  `json:"heater_flame"`
				HeaterTemperature int   `json:"heater_temperature"`
				HeaterErrors      []any `json:"heater_errors"`
			} `json:"heater"`
			EngineRemains          int  `json:"engine_remains"`
			CANClimate             bool `json:"CAN_climate"`
			CANClimateAc           bool `json:"CAN_climate_ac"`
			CANClimateSteeringHeat bool `json:"CAN_climate_steering_heat"`
			CANClimateGlassHeat    bool `json:"CAN_climate_glass_heat"`
			CANClimateSeatVentLvl  int  `json:"CAN_climate_seat_vent_lvl"`
			CANClimateSeatHeatLvl  int  `json:"CAN_climate_seat_heat_lvl"`
			CANClimateEvbHeat      bool `json:"CAN_climate_evb_heat"`
			CANClimateDefroster    bool `json:"CAN_climate_defroster"`
			CANClimateTemp         int  `json:"CAN_climate_temp"`
		} `json:"1079312388"`
	} `json:"stats"`
	EcEventsUpdate bool `json:"ec_events_update"`
	Ucr            struct {
	} `json:"ucr"`
}
