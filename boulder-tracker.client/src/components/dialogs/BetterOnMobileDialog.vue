<template>
  <v-dialog max-width="50vw" ref="dialog" v-model="isVisible">
    <v-card>
      <v-card-title>
        <div class="text-h4">Diese Seite ist für Mobilgeräte optimiert</div>
      </v-card-title>
      <v-card-text class="d-flex flex-column gap-2rem align-center">
        <div class="text-body-1">Tut mir leid, aber diese Seite ist für Mobilgeräte optimiert, weswegen du auf deinem Gerät Schwierigkeiten bekommen könntest
          Du kannst die Anwendung trotzdem weiter nutzen oder den folgenden QR-Code abscannen um die Seite mit deinem Handy zu benutzen.<br/>
        </div>
        <v-img :width="150"
                :aspect-ratio="1"
                cover
                src="src/assets/qr-code.png">
        </v-img>
      </v-card-text>
      <template v-slot:actions>
        <v-btn
            text="Alles klar"
            @click="handleClick"
        ></v-btn>
      </template>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import {onMounted, Ref, ref} from "vue";

const props = defineProps<{
  isVisible: boolean
}>();

const isVisible: Ref<boolean> = ref(false);

onMounted(() => {
  isVisible.value = props.isVisible;
})

function handleClick() {
  isVisible.value = false;
}

</script>

<style scoped>

</style>