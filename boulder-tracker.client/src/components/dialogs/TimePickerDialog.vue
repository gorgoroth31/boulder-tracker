<template>
  <v-dialog v-model="open">
    <template v-slot:activator="{ props: activatorProps }">
      <div class="d-flex flex-column">
        <v-label>{{ label }}</v-label>
        <v-btn v-bind="activatorProps" @click="onOpenDialog">
          {{ props.dateTime.toTimeString().substring(0,5) }}
        </v-btn>
      </div>
    </template>
    <v-time-picker title="" format="24hr" v-model="dateTimeRef">
      <template v-slot:actions>
        <v-btn @click="closeWithoutSend">Abbrechen</v-btn>
        <v-btn @click="submit">Ok</v-btn>
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

const emit = defineEmits<{
  (e: 'submit', success: boolean, hours: number, minutes: number): void
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
    
    dateTimeRef.value = new Date();
    dateTimeRef.value.setHours(hour.value, minute.value)
  }
})

onMounted(() => {
  resetDialog();
})

function resetDialog() {
  hour.value = props.dateTime.getHours()
  minute.value = props.dateTime.getMinutes()
  dateTimeRef.value = new Date();
  dateTimeRef.value.setHours(hour.value, minute.value)
}

function onOpenDialog() {
  resetDialog();
}

function closeWithoutSend() {
  emit('submit', false, 0, 0)
  open.value = false;
}

function submit() {
  emit('submit', true, hour.value, minute.value)
  closeWithoutSend();
}


</script>