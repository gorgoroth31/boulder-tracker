<template>
  <v-dialog>
    <template v-slot:activator="{ props: activatorProps }">
      <v-sheet @click="resetDialog" v-bind="activatorProps" v-ripple class="d-flex justify-center align-center py-2" width="100%">
        <v-icon size="3rem" color="surfaceLighter">mdi-plus-circle-outline</v-icon>
      </v-sheet>
    </template>
    <v-card>
      <v-card-title>Route hinzufügen</v-card-title>
      <v-card-text class="d-flex flex-column gap-1rem">
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
            <v-icon-btn :disabled="boulder.sessionsTried <= 1" @click="boulder.sessionsTried--" icon="mdi-minus"></v-icon-btn>
            <span>{{ boulder.sessionsTried }}</span>
            <v-icon-btn :disabled="boulder.sessionsTried >= 99" @click="boulder.sessionsTried++" icon="mdi-plus"></v-icon-btn>
          </div>
        </v-row>

        <v-row class="d-flex flex-column">
          <v-label>War die Route anstrengend?</v-label>
          <v-checkbox-btn class="align-self-center" @click="boulder.exhausting"></v-checkbox-btn>
        </v-row>
        
        <v-row class="d-flex flex-column">
          <v-label>Verdient diese Route ein Like?</v-label>
          <v-icon class="align-self-center" @click="boulder.like = !boulder.like" :color="boulder.like ? 'red' : 'surfaceLighter'" size="3rem">mdi-heart-outline</v-icon>
        </v-row>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import {onMounted, ref, Ref} from "vue";
import {Boulder} from "@/models/boulder";
import {Difficulty} from "@/models/difficulty";

const boulder: Ref<Boulder> = ref({});

const allDifficulties: Ref<Difficulty[]> = ref([]);

const props = defineProps<{}>()

const emit = defineEmits({
  submit(success: boolean, routeToAdd: Boulder) {
  }
})

onMounted(() => {
  // load difficulties for user from api
  resetDialog();
})

function resetDialog() {
  boulder.value = {
    attempts: 1,
    sessionsTried: 1, 
  }
}

function closeWithoutSend() {
  emit('submit', false, boulder.value)
}

function submit() {
  closeWithoutSend();
}

</script>