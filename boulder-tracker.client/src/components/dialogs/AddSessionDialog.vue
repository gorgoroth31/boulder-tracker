<template>
    <v-dialog :width="width" :height="height * 0.83">
        <template v-slot:activator="{ props: activatorProps }">
            <v-btn :class="btnClass" v-bind="activatorProps" color="green-darken-1" text="Neue Session"
                variant="flat"></v-btn>
        </template>

        <template v-slot:default="{ isActive }">
            <v-card title="Session hinzufügen" class="h-100">
                <v-form v-model="valid" class="h-100 d-flex flex-column justify-content-between">
                    <div class="flex-fill">
                        <div class="d-flex justify-content-center" v-if="currentStep === 1">
                            <v-date-picker :width="width" weekday-format="short" first-day-of-week="1" header=""
                                :max="new Date()" title="Wann warst du bouldern?" show-adjacent-months
                                v-bind:model-value="date"></v-date-picker>
                        </div>

                        <div class="d-flex justify-content-center" v-if="currentStep === 2">
                            <v-time-picker v-model="from" :width="width" format="24hr" title="Von"></v-time-picker>
                        </div>

                        <div class="d-flex justify-content-center" v-if="currentStep === 3">
                            <v-time-picker v-model="to" :width="width" format="24hr" title="Bis"></v-time-picker>
                        </div>
                    </div>

                    <v-stepper-actions :prev-text="prevText" :next-text="nextText"
                        @click:prev="() => handlePrev(isActive)" @click:next="() => handleNext(isActive)">
                    </v-stepper-actions>
                </v-form>
            </v-card>
        </template>
    </v-dialog>
</template>

<script setup lang="ts">

import { computed, onMounted, Ref, ref } from "vue";
import { Session } from "./../../models/session"
import { DateRange } from "../../models/daterange";

import { useDisplay } from 'vuetify'

const { mobile, width, height } = useDisplay()

const props = defineProps({
    btnClass: String
})

const prevText = computed(() => {
    return currentStep.value === 1 ? "Verlassen" : "Zurück"
})

const nextText = computed(() => {
    return currentStep.value === maxSteps ? "Speichern" : "Weiter"
})

// todo add validation
// only can go forward if value is provided
// for time, only can select time, which is AFTER from

let valid: boolean = false;
let date: Date = new Date()
let from: Date = new Date()
let to: Date = new Date()

const currentStep: Ref<number> = ref(1);

const maxSteps: number = 4;

onMounted(() => {
    session.value.visitTime = new DateRange
    if (mobile.value) {

    }
})

function handlePrev(isActive: Ref<boolean, boolean>) {
    if (currentStep.value === 1) {
        isActive.value = false
        return;
    }

    currentStep.value--;
}

function handleNext(isActive: Ref<boolean, boolean>) {
    if (currentStep.value === maxSteps) {
        isActive.value = false
        return;
    }

    currentStep.value++;
}

const session = ref(new Session())
</script>