// RGB Band
Adafruit_NeoPixel strip = Adafruit_NeoPixel(15, RGB_PIN, NEO_GRB + NEO_KHZ800);

bool bandSwap = false;
int bandCpt = 0;
int bandCurrentAction;

void startBand(){
  // RGB Band
  strip.begin();
  strip.show(); // Initialize all pixels to 'off'
}

void setBandAction(byte action){
  bandCurrentAction = action - 48;
}

void doBandShortAction(){
   switch (bandCurrentAction) {
    case 2:
      doK2000Light();
      break;
  }
}

void doBandAction(){
   switch (bandCurrentAction) {
    case 1:
      doPoliceLight();
      break;
    case 0:
      startBand();
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
    bandCpt = bandCpt+1;
  }
  strip.show();
}

