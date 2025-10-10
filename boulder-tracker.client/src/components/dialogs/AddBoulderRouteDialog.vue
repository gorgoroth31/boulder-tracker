<template>
  <v-dialog v-model="isOpen">
    <template v-slot:activator="{ props: activatorProps }">
      <v-sheet @click="resetDialog" v-bind="activatorProps" v-ripple class="d-flex justify-center align-center py-2"
               width="100%">
        <v-icon size="3rem" color="surfaceLighter">mdi-plus-circle-outline</v-icon>
      </v-sheet>
    </template>
    <v-card>
      <v-card-title>Route hinzufügen</v-card-title>
      <v-card-text class="d-flex flex-column gap-1rem">
        <v-row class="d-flex flex-column">
          <v-select label="Geschraubte Schwierigkeit" v-model="boulder.screwedDifficulty" :items="allDifficulties"
                    return-object
                    item-title="alias">
          </v-select>
        </v-row>

        <v-row class="d-flex flex-column">
          <v-select label="Gefühlte Schwierigkeit" v-model="boulder.feltLikeDifficulty" :items="allDifficulties"
                    return-object
                    item-title="alias">
          </v-select>
        </v-row>

        <v-row class="d-flex flex-column">
          <v-select
              v-model="boulder.style"
              :items="allStyles"
              label="Styles auswählen"
              return-object
              item-title="alias"
              multiple>
            <template v-slot:selection="{ item, index }">
              <v-chip :text="item.title"></v-chip>
            </template>
          </v-select>
        </v-row>

        <v-row class="d-flex flex-column">
          <v-label>Anzahl der Versuche</v-label>
          <div class="d-flex align-center gap-1rem">
            <v-icon-btn :disabled="boulder.attempts <= 1" @click="boulder.attempts--" icon="mdi-minus"></v-icon-btn>
            <span>{{ boulder.attempts }}</span>
            <v-icon-btn :disabled="boulder.attempts >= 99" @click="boulder.attempts++" icon="mdi-plus"></v-icon-btn>
          </div>
        </v-row>

        <v-row class="d-flex flex-column">
          <v-label>Anzahl benötigter Sessions</v-label>
          <div class="d-flex align-center gap-1rem">
            <v-icon-btn :disabled="boulder.sessionsTried <= 1" @click="boulder.sessionsTried--"
                        icon="mdi-minus"></v-icon-btn>
            <span>{{ boulder.sessionsTried }}</span>
            <v-icon-btn :disabled="boulder.sessionsTried >= 99" @click="boulder.sessionsTried++"
                        icon="mdi-plus"></v-icon-btn>
          </div>
        </v-row>

        <v-row class="d-flex flex-column">
          <v-label>War die Route anstrengend?</v-label>
          <v-checkbox-btn class="align-self-center" @click="boulder.exhausting = !boulder.exhausting"></v-checkbox-btn>
        </v-row>

        <v-row class="d-flex flex-column">
          <v-label>Verdient diese Route ein Like?</v-label>
          <v-icon class="align-self-center" @click="boulder.like = !boulder.like"
                  :color="boulder.like ? 'red' : 'surfaceLighter'" size="3rem">mdi-heart-outline
          </v-icon>
        </v-row>
        <v-row v-if="errorMessage">
          <v-alert color="error">{{ errorMessage }}</v-alert>
        </v-row>
      </v-card-text>
      <template v-slot:actions>
        <v-btn text="Abbrechen"
               @click="handleCancel"
        ></v-btn>
        <v-btn text="Hinzufügen"
               @click="handleSubmit"
        ></v-btn>
      </template>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import {onMounted, ref, Ref} from "vue";
import {Boulder} from "@/models/boulder";
import {Difficulty} from "@/models/difficulty";
import {getAllDifficulties} from "@/api/difficulty.api";
import {Style} from "node:util";
import {getAllStyles} from "@/api/styles.api";

const isLoading: Ref<boolean> = ref(false);
const isOpen: Ref<boolean> = ref(false);

const boulder: Ref<Boulder> = ref({});

const allDifficulties: Ref<Difficulty[]> = ref([]);
const allStyles: Ref<Style[]> = ref([]);

const errorMessage: Ref<string> = ref("");

const emit = defineEmits<{
  (e: 'submit', success: boolean, routeToAdd: Boulder): void
}>()

onMounted(() => {
  isLoading.value = true;
  getAllDifficulties().then(result => {
    allDifficulties.value = result.data
    allDifficulties.value.sort((a: Difficulty, b: Difficulty) => a.relativeLevel - b.relativeLevel);
  })

  getAllStyles().then(result => {
    allStyles.value = result.data
  })
  resetDialog();
})

function resetDialog() {
  isLoading.value = true;
  boulder.value = {
    attempts: 1,
    sessionsTried: 1,
    like: false,
    exhausting: false,
    style: []
  }
  isLoading.value = false;
}

function handleCancel() {
  emit('submit', false, null);
  isOpen.value = false;
}

function handleSubmit() {
  const isBoulderValid = checkBoulderValidation();

  if (!isBoulderValid) {
    errorMessage.value = "Es fehlen noch einige Daten, bitte checke nochmal alle Felder."
    return;
  }

  errorMessage.value = "";

  emit('submit', true, boulder.value)
  isOpen.value = false;
}

function checkBoulderValidation(): boolean {
  return boulder.value.feltLikeDifficulty !== undefined && boulder.value.screwedDifficulty !== undefined;
}

</script>