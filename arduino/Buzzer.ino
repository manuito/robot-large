// Mario main theme melody
int mario_melody[] = {
  NOTE_E7, NOTE_E7, 0, NOTE_E7,
  0, NOTE_C7, NOTE_E7, 0,
  NOTE_G7, 0, 0,  0,
  NOTE_G6, 0, 0, 0,
 
  NOTE_C7, 0, 0, NOTE_G6,
  0, 0, NOTE_E6, 0,
  0, NOTE_A6, 0, NOTE_B6,
  0, NOTE_AS6, NOTE_A6, 0,
 
  NOTE_G6, NOTE_E7, NOTE_G7,
  NOTE_A7, 0, NOTE_F7, NOTE_G7,
  0, NOTE_E7, 0, NOTE_C7,
  NOTE_D7, NOTE_B6, 0, 0,
 
  NOTE_C7, 0, 0, NOTE_G6,
  0, 0, NOTE_E6, 0,
  0, NOTE_A6, 0, NOTE_B6,
  0, NOTE_AS6, NOTE_A6, 0,
 
  NOTE_G6, NOTE_E7, NOTE_G7,
  NOTE_A7, 0, NOTE_F7, NOTE_G7,
  0, NOTE_E7, 0, NOTE_C7,
  NOTE_D7, NOTE_B6, 0, 0
};

// Mario main them tempo
int mario_tempo[] = {
  12, 12, 12, 12,
  12, 12, 12, 12,
  12, 12, 12, 12,
  12, 12, 12, 12,
 
  12, 12, 12, 12,
  12, 12, 12, 12,
  12, 12, 12, 12,
  12, 12, 12, 12,
 
  9, 9, 9,
  12, 12, 12, 12,
  12, 12, 12, 12,
  12, 12, 12, 12,
 
  12, 12, 12, 12,
  12, 12, 12, 12,
  12, 12, 12, 12,
  12, 12, 12, 12,
 
  9, 9, 9,
  12, 12, 12, 12,
  12, 12, 12, 12,
  12, 12, 12, 12,
};

// Underworld melody
int underworld_melody[] = {
  NOTE_C4, NOTE_C5, NOTE_A3, NOTE_A4,
  NOTE_AS3, NOTE_AS4, 0,
  0,
  NOTE_C4, NOTE_C5, NOTE_A3, NOTE_A4,
  NOTE_AS3, NOTE_AS4, 0,
  0,
  NOTE_F3, NOTE_F4, NOTE_D3, NOTE_D4,
  NOTE_DS3, NOTE_DS4, 0,
  0,
  NOTE_F3, NOTE_F4, NOTE_D3, NOTE_D4,
  NOTE_DS3, NOTE_DS4, 0,
  0, NOTE_DS4, NOTE_CS4, NOTE_D4,
  NOTE_CS4, NOTE_DS4,
  NOTE_DS4, NOTE_GS3,
  NOTE_G3, NOTE_CS4,
  NOTE_C4, NOTE_FS4, NOTE_F4, NOTE_E3, NOTE_AS4, NOTE_A4,
  NOTE_GS4, NOTE_DS4, NOTE_B3,
  NOTE_AS3, NOTE_A3, NOTE_GS3,
  0, 0, 0
};

// Underwolrd tempo
int underworld_tempo[] = {
  12, 12, 12, 12,
  12, 12, 6,
  3,
  12, 12, 12, 12,
  12, 12, 6,
  3,
  12, 12, 12, 12,
  12, 12, 6,
  3,
  12, 12, 12, 12,
  12, 12, 6,
  6, 18, 18, 18,
  6, 6,
  6, 6,
  6, 6,
  18, 18, 18, 18, 18, 18,
  10, 10, 10,
  10, 10, 10,
  3, 3, 3
};

int buzzerCpt = 0;
int buzzerCurrentAction;

void setBuzzerAction(byte action){
  buzzerCurrentAction = action - 48;
}

void doBuzzerAction(){
   switch (buzzerCurrentAction) {
    case 1:
      doBeep();
      buzzerCurrentAction = 0;
      break;
    case 0:
      break;
  }
}

void doBuzzerShortAction(){
   switch (buzzerCurrentAction) {
    case 2:
      doSingMarioTheme();
      break;
    case 3:
      doSingUnderworldTheme();
      break;
  }
}

void doBeep(){
   tone(BUZZER_PIN, 600, 200);
}

void doSingMarioTheme() {
  sing(1);
}

void doSingUnderworldTheme() {
  sing(2);
}

int getBuzzerAction(){
  return buzzerCurrentAction;
}

void sing(int song) {
  // iterate over the notes of the melody:
  if (song == 2) {
    processSong(underworld_melody, underworld_tempo, 56);
  } else { 
    processSong(mario_melody, mario_tempo, 78);
  }
}

void processSong(int melody[], int tempo[], int size) {

  if(buzzerCpt < size){
      // Stop previous note if any
      
      // to calculate the note duration, take one second
      // divided by the note type.
      //e.g. quarter note = 1000 / 4, eighth note = 1000/8, etc.
      int noteDuration = 1000 / tempo[buzzerCpt];
      float note = melody[buzzerCpt];

      Serial.print("Expected Note : ");
      if(note > 0){
      Serial.print("Note : ");
      
      Serial.print(buzzerCpt);
      Serial.print("/");
      Serial.print(size);
      Serial.print(" => ");
      Serial.print(note);
      Serial.print(" for ");
      Serial.print(tempo[buzzerCpt]);
      Serial.println("");
      tone(BUZZER_PIN, note, noteDuration);
 
      // to distinguish the notes, set a minimum time between them.
      // the note's duration + 30% seems to work well:
    //  int pauseBetweenNotes = noteDuration * 1.30;
    //  delay(pauseBetweenNotes);
      } else {
        tone(BUZZER_PIN, 0, noteDuration);
      }
      // stop the tone playing:
      buzzerCpt++;
  } else {
    buzzerCpt = 0;
    buzzerCurrentAction = 0; // Stop when completed
  }
}
