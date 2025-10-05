<template>
  <v-progress-circular v-if="isLoading"></v-progress-circular>
  <v-form v-else>
    <v-container>
      <v-row>
        <v-date-input label="Wann warst du bouldern?"
                      :max="new Date()"
                      first-day-of-week="1"
                      show-adjacent-months
                      weekday-format="short"
                      display-format="normalDate"
                      autocomplete="off"
                      v-model="sessionDate">
        </v-date-input>
      </v-row>
      <v-row>
        <time-picker-dialog label="Von:" :dateTime="session.startTime"></time-picker-dialog>
        <v-spacer></v-spacer>
        <time-picker-dialog label="Bis:" :dateTime="session.endTime"></time-picker-dialog>
        <v-spacer></v-spacer>
      </v-row>
      <v-row>
        <v-btn @click="save">Speichern</v-btn>
      </v-row>
    </v-container>
  </v-form>
</template>

<script setup lang="ts">
import {onMounted, Ref, ref, watch} from "vue";
import mainPageUtils from "./../utils/mainPageUtils";
import {getCurrentInProgressSession, updateSession} from "@/api/session.api";
import {Session} from "@/models/session";
import TimePickerDialog from "@/components/dialogs/TimePickerDialog.vue";

const isLoading: Ref<boolean> = ref(true);

const session: Ref<Session> = ref<Session>({});
const sessionDate: Ref<Date> = ref(new Date());

watch(sessionDate, (newVal, _) => {
  const oldStart = new Date(session.value.startTime);
  const oldEnd = new Date(session.value.endTime);

  const newStart = new Date(newVal);
  const newEnd = new Date(newVal);

  newStart.setHours(oldStart.getHours(), oldStart.getMinutes());
  newEnd.setHours(oldEnd.getHours(), oldEnd.getMinutes());

  session.value.startTime = newStart;
  session.value.endTime = newEnd;
})

onMounted(() => {
  isLoading.value = true;
  mainPageUtils.pageTitle.value = "Session hinzufügen";

  getCurrentInProgressSession().then(value => {
    session.value = value.data;
    sessionDate.value = session.value.startTime;
    isLoading.value = false;
  })
})

function save() {
  updateSession(session.value)
}
</script>