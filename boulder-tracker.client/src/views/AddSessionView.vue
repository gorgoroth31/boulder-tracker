<template>
  <v-progress-linear indeterminate v-if="isLoading"></v-progress-linear>
  <v-form v-else>
    <v-container class="d-flex flex-column gap-1rem">
      <v-row class="d-flex flex-column">
        <v-date-input label="Wann warst du bouldern?"
                      :max="new Date()"
                      first-day-of-week="1"
                      show-adjacent-months
                      weekday-format="short"
                      display-format="normalDate"
                      autocomplete="off"
                      v-model="sessionDate">
        </v-date-input>
        <time-picker-dialog @submit="setNewStartTime" label="Von:" :dateTime="new Date(session.startTime)"></time-picker-dialog>
        <time-picker-dialog @submit="setNewEndTime" label="Bis:" :dateTime="new Date(session.endTime)"></time-picker-dialog>
        <v-snackbar v-model="showEndtimeSnackbar" color="info" timeout="3000">
          Bitte beachte, dass deine Endzeit nicht vor der Startzeit liegen kann
        </v-snackbar>
      </v-row>
      <v-row class="d-flex flex-column gap-1rem">
        <div class="text-h5">Routen</div>
        <v-expansion-panels>
          <BoulderRouteCard v-for="boulder in session.routesSolved"
                            :boulderRoute="boulder"
                            @delete="handleBoulderDelete">
          </BoulderRouteCard>
        </v-expansion-panels>

        <AddBoulderRouteCardDialog @submit="handleAddBoulderRouteCallback"></AddBoulderRouteCardDialog>
      </v-row>
      <v-row class="d-flex justify-end">
        <v-btn ref="submit-btn" :disabled="!isSaved">Session abschließen</v-btn>
        
        <ConfirmationDialog @submit="submit" :title="null" :activator="submitButtonRef" text="Bitte bestätige, dass du die Session abschließen möchtest" ok-button-text="Ok" cancel-button-text="Abbrechen" :close-on-back="true"></ConfirmationDialog>
      </v-row>
    </v-container>
  </v-form>
</template>

<script setup lang="ts">
import {onMounted, Ref, ref, useTemplateRef, watch} from "vue";
import mainPageUtils from "./../utils/mainPageUtils";
import {getCurrentInProgressSession, submitCurrentSession, updateSession} from "@/api/session.api";
import {Session} from "@/models/session";
import TimePickerDialog from "@/components/dialogs/TimePickerDialog.vue";
import BoulderRouteCard from "@/components/BoulderRouteCard.vue";
import AddBoulderRouteCardDialog from "@/components/dialogs/AddBoulderRouteDialog.vue";
import {Boulder} from "@/models/boulder";
import {removeItem} from "@/utils/arrayUtils";
import {deleteBoulder} from "@/api/boulder.api";
import router from "@/router";
import { useDebounceFn } from '@vueuse/core'
import ConfirmationDialog from "@/components/dialogs/ConfirmationDialog.vue";
import {sleep} from "@/utils/otherUtils";

const submitButtonRef = useTemplateRef("submit-btn")

const isLoading: Ref<boolean> = ref(true);
const showEndtimeSnackbar: Ref<boolean> = ref(false);

const session: Ref<Session> = ref<Session>({});
const sessionDate: Ref<Date> = ref(new Date());

const isSaved: Ref<boolean> = ref(false);

watch(sessionDate, (newVal, _) => {
  const oldStart = new Date(session.value.startTime);
  const oldEnd = new Date(session.value.endTime);

  const newStart = new Date(newVal);
  const newEnd = new Date(newVal);

  newStart.setHours(oldStart.getHours(), oldStart.getMinutes());
  newEnd.setHours(oldEnd.getHours(), oldEnd.getMinutes());

  session.value.startTime = new Date(newStart);
  session.value.endTime = new Date(newEnd);
  saveWithDebounce();  
})

onMounted(() => {
  isLoading.value = true;
  mainPageUtils.pageTitle.value = "Session hinzufügen";

  getCurrentInProgressSession().then(value => {
    session.value = value.data;
    parseDate();
    sessionDate.value = session.value.startTime;
    isLoading.value = false;
  })
})

const saveWithDebounce = useDebounceFn(async () => {
  isSaved.value = false;
  await updateSession(session.value).then(value => {
    session.value = value.data;
  })
  parseDate();
  await sleep(666);
  isSaved.value = true;
}, 400)

function parseDate() {
  session.value.startTime = new Date(session.value.startTime)
  session.value.endTime = new Date(session.value.endTime)
}

function setNewStartTime(success: boolean, hours: number, minutes: number) {
  if (!success) {
    return;
  }
  session.value.startTime.setHours(hours, minutes);
  saveWithDebounce();
}

function setNewEndTime(success: boolean, hours: number, minutes: number) {
  if (!success) {
    return;
  }
  const endTimeBackup = new Date(session.value.endTime);
  session.value.endTime.setHours(hours, minutes);
  
  if (session.value.endTime < session.value.startTime) {
    showEndtimeSnackbar.value = true;
    session.value.endTime = endTimeBackup;
    return
  }
  
  session.value.endTime.setHours(hours, minutes);
  saveWithDebounce();
}

function handleBoulderDelete(route: Boulder) {
  if (route.id === undefined) {
    session.value.routesSolved = removeItem<Boulder>(session.value.routesSolved, route)
    return
  }

  deleteBoulder(route.id).then(value => {
    if (value.status === 204) {
      session.value.routesSolved = removeItem<Boulder>(session.value.routesSolved, route)
    }
  })
}

function handleAddBoulderRouteCallback(success: boolean, routeToAdd: Boulder) {
  if (!success) {
    return;
  }

  routeToAdd.sessionId = session.value.id;
  session.value.routesSolved.push(routeToAdd);
  saveWithDebounce()
}

async function submit(isSuccess: boolean) {
  if (!isSuccess) {
    return;
  }
  
  await saveWithDebounce()
  submitCurrentSession().then(value => {
    if (value.status === 204) {
      router.push("/")
    }
  })
}
</script>