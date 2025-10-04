<template>
  <div class="text-h5 gy-0">Wann warst du bouldern?</div>
  <v-form>
    <v-container>
      <v-row>
        <DatePickerDialog title="Wann war die Session?" 
                          :width="width" 
                          header=""
                          :max="new Date()"
                          :height="height * 0.83"
                          :date="session."
                          >
        </DatePickerDialog>
      </v-row>
    </v-container>
  </v-form>
</template>

<script setup lang="ts">
import {onMounted, Ref, ref} from "vue";
  import mainPageUtils from "./../utils/mainPageUtils";
  import {Session} from "node:inspector";
import {getCurrentInProgressSession} from "@/api/session.api";
import DatePickerDialog from "@/components/dialogs/DatePickerDialog.vue";
import {useDisplay} from "vuetify/framework";

  const session: Ref<Session> = ref<Session>({});

  const { width, height } = useDisplay()
  
  onMounted(() => {
    mainPageUtils.pageTitle.value = "Session hinzufügen";
    
    getCurrentInProgressSession().then(value => {
      session.value = value.data;
      console.log(session.value);
    })
  })
</script>