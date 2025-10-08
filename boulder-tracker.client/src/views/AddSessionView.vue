<template>
  <v-progress-circular v-if="isLoading"></v-progress-circular>
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
        <time-picker-dialog label="Von:" :dateTime="session.startTime"></time-picker-dialog>
        <time-picker-dialog label="Bis:" :dateTime="session.endTime"></time-picker-dialog>
      </v-row>
      <v-row>
        <v-checkbox v-model="session.boulderedSolo" color="primary" label="Warst du alleine bouldern?"></v-checkbox>
      </v-row>
      <v-row>
        <div class="text-h5">Routen</div>
        <!--First display all boulder routes as cards, then a card with a +, when clicking +, dialog opens to add boulder-->
      </v-row>
      <v-row class="d-flex justify-end">
        <v-btn @click="save" class="border-right-0">Speichern</v-btn>
        <v-btn
            class="border-left-0"
            id="menu-activator">
          <v-icon>mdi-chevron-down</v-icon>
        </v-btn>

        <v-menu activator="#menu-activator">
          <v-list class="pa-0">
            <v-list-item class="pa-0">
              <v-list-item-title class="pa-0">
                <v-btn @click="submit" class="border-right-0">Speichern & Abschließen</v-btn>
              </v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
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
import BoulderRouteCard from "@/components/BoulderRouteCard.vue";

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

function submit() {
  alert("implement submit with confirmation dialog")
}

function save() {
  updateSession(session.value)
}
</script>