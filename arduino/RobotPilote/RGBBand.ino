// RGB Band
Adafruit_NeoPixel strip = Adafruit_NeoPixel(RGB_LED_COUNT, RGB_PIN, NEO_GRB + NEO_KHZ800);

bool bandSwap = false;
int bandCpt = 0;
uint8_t bandCurrentAction;

void startBand(){
  // RGB Band
  strip.begin();
  strip.show(); // Initialize all pixels to 'off'
}

void setBandAction(byte action){
  bandCurrentAction = action - 48;
  doOff();
  bandCpt = 0;
}

uint8_t getBandAction(){
  return bandCurrentAction;
}

void doBandShortAction(){
   switch (bandCurrentAction) {
    case 2:
      doK2000Light();
      break;
    case 3:
      doRainbowCycle();
      break;
    case 4:
      doRadar();
      break;
  }
}

void doBandAction(){
   switch (bandCurrentAction) {
    case 1:
      doPoliceLight();
      break;
    case 0:
      doOff();
      break;
  }
}

void doPoliceLight(){
  if(bandSwap){
    strip.setPixelColor(5, strip.Color(0, 100, 0));
    strip.setPixelColor(10, strip.Color(100, 0, 0));
  } else {
    strip.setPixelColor(5, strip.Color(100, 0, 0));
    strip.setPixelColor(10, strip.Color(0, 100, 0));
  }
  strip.show();
  bandSwap = !bandSwap;
}

void doK2000Light(){

  if(bandCpt <= 4){
    strip.setPixelColor(4 + bandCpt, strip.Color(20, 0, 0));
    strip.setPixelColor(5 + bandCpt, strip.Color(60, 0, 0));
    strip.setPixelColor(6 + bandCpt, strip.Color(90, 0, 0));
    strip.setPixelColor(7 + bandCpt, strip.Color(100, 0, 0));
  } else {
    strip.setPixelColor(15 - bandCpt, strip.Color(20, 0, 0));
    strip.setPixelColor(14 - bandCpt, strip.Color(60, 0, 0));
    strip.setPixelColor(13 - bandCpt, strip.Color(90, 0, 0));
    strip.setPixelColor(12 - bandCpt, strip.Color(100, 0, 0));
  }

  if(bandCpt >= 8){
    bandCpt = 0;
  } else {
    bandCpt++;
  }
  strip.show();
}

void doRainbowCycle() {
  
  if(bandCpt >= 256*5){
    bandCpt = 0;
  } else {
    bandCpt++;
  }
  
  for(uint8_t i=0; i< RGB_LED_COUNT; i++) {
    strip.setPixelColor(i, Wheel(((i * 256 / RGB_LED_COUNT) + bandCpt) & 255));
  }
  strip.show();
}

void doOff(){
  for(int i = 0; i < RGB_LED_COUNT; i++){
    strip.setPixelColor(i, strip.Color(0, 0, 0));
  }
  strip.show();
}

void doRadar(){

  if(getLeftDist() < 20){
    if(getLeftDist() < 10){
      strip.setPixelColor(8, strip.Color(150, 0, 0));
      strip.setPixelColor(9, strip.Color(150, 0, 0));
    } else {
      strip.setPixelColor(8, strip.Color(50, 0, 0));
      strip.setPixelColor(9, strip.Color(50, 0, 0));
    }
  } else {
    strip.setPixelColor(8, strip.Color(0, 0, 0));
    strip.setPixelColor(9, strip.Color(0, 0, 0));
  }
  
  if(getRightDist() < 20){
    if(getRightDist() < 10){
      strip.setPixelColor(6, strip.Color(150, 0, 0));
      strip.setPixelColor(7, strip.Color(150, 0, 0));
    } else {
      strip.setPixelColor(6, strip.Color(50, 0, 0));
      strip.setPixelColor(7, strip.Color(50, 0, 0));
    }
  } else {
    strip.setPixelColor(6, strip.Color(0, 0, 0));
    strip.setPixelColor(7, strip.Color(0, 0, 0));
  }
  
  strip.show();
}

// Input a value 0 to 255 to get a color value.
// The colours are a transition r - g - b - back to r.
uint32_t Wheel(byte WheelPos) {
  WheelPos = 255 - WheelPos;
  if(WheelPos < 85) {
    return strip.Color(255 - WheelPos * 3, 0, WheelPos * 3);
  }
  if(WheelPos < 170) {
    WheelPos -= 85;
    return strip.Color(0, WheelPos * 3, 255 - WheelPos * 3);
  }
  WheelPos -= 170;
  return strip.Color(WheelPos * 3, 255 - WheelPos * 3, 0);
}
