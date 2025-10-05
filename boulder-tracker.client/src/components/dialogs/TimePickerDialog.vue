<template>
  <v-dialog v-model="open">
    <template v-slot:activator="{ props: activatorProps }">
      <div class="d-flex flex-column">
        <v-label>{{ label }}</v-label>
        <v-btn v-bind="activatorProps">
          {{ dateTimeRef.toTimeString().substring(0,5) }}
        </v-btn>
      </div>
    </template>
    <v-time-picker title="" format="24hr" v-model="dateTimeRef">
      <template v-slot:actions>
        <v-btn @click="open = false">Schließen</v-btn>
      </template>
    </v-time-picker>
  </v-dialog>
</template>

<script setup lang="ts">
import {onMounted, ref, Ref, watch} from "vue";

const props = defineProps<{
  dateTime: Date,
  label: string
}>()

const open: Ref<boolean> = ref(false);

const dateTimeRef: Ref<Date> = ref(new Date());

const hour: Ref<number> = ref(0);
const minute: Ref<number> = ref(0);

watch(dateTimeRef, (val) => {
  const timeRegex = new RegExp("^([01]?[0-9]|2[0-3]):[0-5][0-9]$");
  
  if (typeof val === "string" && timeRegex.test(val)) {
    const newVal = val as string;
    hour.value = Number(newVal.substring(0, 2))
    minute.value = Number(newVal.substring(3, 5))
    
    props.dateTime.setHours(hour.value, minute.value)
    
    dateTimeRef.value = props.dateTime;
  }
})

onMounted(() => {
  dateTimeRef.value = props.dateTime;
  hour.value = props.dateTime.getHours()
  minute.value = props.dateTime.getMinutes()
})
</script>