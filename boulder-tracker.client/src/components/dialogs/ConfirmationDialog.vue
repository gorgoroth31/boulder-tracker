<template>
  <v-dialog :activator="activator" max-width="100vw" :close-on-back="closeOnBack" ref="dialog" v-model="isVisible">
      <v-card
          :text="text"
          :title="title">
        <template v-slot:actions>
          <v-btn
              :text="cancelButtonText"
              @click="handleCancel"
          ></v-btn>
          <v-btn
              :text="okButtonText"
              @click="handleSubmit"
          ></v-btn>
        </template>
      </v-card>
  </v-dialog>
</template>

<script setup lang="ts">

import {ComponentPublicInstance, CreateComponentPublicInstanceWithMixins, Ref, ref, ShallowRef} from "vue";

defineProps<{
  activator: Element | (string & {}) | 'parent' | ComponentPublicInstance,
  title: string,
  text: string,
  okButtonText: string,
  cancelButtonText: string,
  closeOnBack: boolean,
}>();

const emit = defineEmits({
  submit(success: boolean) {}
})

const isVisible : Ref<boolean, boolean> = ref(false);

function handleSubmit() {
  isVisible.value = false;
  emit('submit', true)
}

function handleCancel() {
  isVisible.value = false;
  emit('submit', false);
}

</script>