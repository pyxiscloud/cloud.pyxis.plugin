syntax = "proto3";
package plugin;
option go_package = "github.com/pyxiscloud/cloud.pyxis.plugin";

message PluginMessage {
  SubtypeOneOf subtype = 1;
  PluginAddress address = 2;
  reserved 3 to max;

}

message Plugin {
  string name = 1;
  bool ableUploadDevices = 2;
  map<string, string> stringValues = 3;
  map<string, int32> numberValues = 4;
  map<string, float> floatValues = 5;

  reserved 6 to max;
}

message SubtypeOneOf {
  oneof subtype {
    float value = 1;
    bool on = 2;
    DIMMER dimmer = 3;
    HSL hsl = 4;
    uint32 hue = 5;
    uint32 saturation = 6;
    RGB rgb = 7;
    uint32 red = 8;
    uint32 green = 9;
    uint32 blue = 10;
    string state = 11;
    double temperature = 12;
    double currentTemp = 13;
    uint32 heatCoolState = 14;
    double targetTemp = 15;
    uint32 batterylevel = 16;
    bool batterylow = 17;
    float rotate = 18;
    string anytext = 19;
    uint32 playStop = 20;
    uint32 volume = 21;
    uint32 operation = 22;
    uint32 fanMode = 23;
    THERMOSTAT thermostat = 24;
    Flip flip = 25;
    uint32 airqualityPm = 26;
    float currentState = 27;
    float targetState = 28;
    bool lockControls = 29;
    uint32 limitValue = 30;
    uint32 chargingState = 31;
    uint32 linkQuality = 32;
    uint32 colorTemp = 33; //140 - 500
    SubtypeOneOf lastStatus = 34;
    bool getListBaseTopics = 35;
    bool setListBaseTopics = 36;
    bool actListBaseTopics = 37;
    uint32 whiteValue = 38;
    uint32 yellowValue = 39;
    Subtype lastStatusSubtype = 40;

    //После добавления надо добавить в BusParser в main.go, SendToPyxisBus
  }

  reserved 41 to max;
}

message Subtype {
  float value = 1;
  bool on = 2;
  DIMMER dimmer = 3;
  HSL hsl = 4;
  uint32 hue = 5;
  uint32 saturation = 6;
  RGB rgb = 7;
  uint32 red = 8;
  uint32 green = 9;
  uint32 blue = 10;
  string state = 11;
  double temperature = 12;
  double currentTemp = 13;
  uint32 heatCoolState = 14;
  double targetTemp = 15;
  uint32 batterylevel = 16;
  bool batterylow = 17;
  float rotate = 18;
  string anytext = 19;
  uint32 playStop = 20;
  uint32 volume = 21;
  uint32 operation = 22;
  uint32 fanMode = 23;
  THERMOSTAT thermostat = 24;
  Flip flip = 25;
  uint32 airqualityPm = 26;
  float currentState = 27;
  float targetState = 28;
  bool lockControls = 29;
  uint32 limitValue = 30;
  uint32 chargingState = 31;
  uint32 linkQuality = 32;
  uint32 colorTemp = 33;
  SubtypeOneOf lastStatus = 34;
  bool getListBaseTopics = 35;
  bool setListBaseTopics = 36;
  bool actListBaseTopics = 37;
  uint32 whiteValue = 38;
  uint32 yellowValue = 39;
  Subtype lastStatusSubtype = 40;
  reserved 41 to max;
}

enum Flip {
  Flip_90 = 0;
  Flip_180 = 1;
  Flip_270 = 2;
  Flip_360 = 3;
  reserved 4 to 360;
}

message DIMMER {
  uint32 value = 1;
  bool on = 2;
  uint32 colorTemp = 3;
  uint32 whiteValue = 4;
  uint32 yellowValue = 5;
  //  uint32 lastBrightnes = 4;
  repeated string subtypeList = 6;
  reserved 7 to max;

  //при добавлении добавлять в main.go SetActualValues
}

message THERMOSTAT {
  double currentTemp = 1;
  uint32 heatCoolState = 2;
  double targetTemp = 3;
  uint32 fanMode = 4;
  repeated string subtypeList = 5;
  reserved 6 to max;
  //при добавлении добавлять в main.go SetActualValues

}

message HSL {
  uint32 hue = 1;
  uint32 saturation = 2;
  uint32 lightness = 3;
  uint32 colorTemp = 4;
  //  uint32 lastBrightnes = 5;
  bool on = 6;
  repeated string subtypeList = 7;
  reserved 8 to max;
  //при добавлении добавлять в main.go SetActualValues

}

message RGB {
  uint32 red = 1;
  uint32 green = 2;
  uint32 blue = 3;
  uint32 alpha = 4;
  bool on = 5;
  repeated string subtypeList = 6;
  reserved 7 to max;
  //при добавлении добавлять в main.go SetActualValues

}

message PluginAddress {
  string name = 1;
  //  map<string, string> mapTypeMap = 2;
  map<string, string> stringValues = 3;
  map<string, int32> numberValues = 4;
  map<string, float> floatValues = 5;
  Type type = 6;

  reserved 7 to max;
}

enum Type {
  LightbulbRelay = 0;
  LightbulbDimmer = 1;
  LightbulbRGB = 2;
  Switch = 3;
  Outlet = 4;
  Control = 5;
  TemperatureSensor = 6;
  WindowCovering = 7;
  ContactSensor = 8;
  LeakSensor = 9;
  SmokeSensor = 10;
  LightSensor = 11;
  HumiditySensor = 12;
  MotionSensor = 13;
  Thermostat = 14;
  FloorHeating = 15;
  MusicPlayer = 16;
  IRemitter = 17;
  AV_control = 18;
  GarageDoorOpener = 19;
  LockMechanism = 20;
  PressureSensor = 21;
  AirQuality = 22;
  AirPurifier = 23;
  Humidifier = 24;
  CarbonDioxideSensor = 25;
  BridgedControl = 26;
  Fan = 27;
  Faucets = 28;
  Sprinklers = 29;
  //После добавления надо добавить в main в main.go, SetActualsValues, SetTargetValues,
  //homekit.ts в runHomekit, во все драйверы, main.ts
  //voice_assistant savedStatuses.go
  //после изменения обновить файл на гитхабе
  reserved 30 to max;
}
