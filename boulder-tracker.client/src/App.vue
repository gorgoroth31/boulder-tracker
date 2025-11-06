<template>
  <v-app :class="'overflow-scroll flex flex-col md:flex-row gap-4 ' + mainPageUtil.appCss.value">
    <div class="w-100 bg-amber text-center">Beta-Version</div>
      <div v-if="isLoading" class="w-100 h-100 d-flex justify-center align-center">
        <v-progress-linear indeterminate></v-progress-linear>
      </div>
      <div class="padding-1rem">
        <div class="text-h4 text-break">{{ mainPageUtil.pageTitle.value }}</div>
        <v-divider thickness="3" class="my-3"></v-divider>
        <div class="mb-15">
          <RouterView/>
        </div>
      </div>
      <NavMenu></NavMenu>
  </v-app>
</template>

<script setup lang="ts">
import { RouterView} from 'vue-router'
import mainPageUtil from "./utils/mainPageUtils";
import NavMenu from "@/components/layout/NavMenu.vue";
import { onMounted, ref, Ref } from 'vue';
import { getCurrentLoggedInUser } from './api/user.api';
import mainPageUtils from './utils/mainPageUtils';

const isLoading: Ref<boolean> = ref(true);

onMounted(async () => {
    await getCurrentLoggedInUser().then(value => {
    mainPageUtils.pageTitle.value = "Hallo " + value.data.userName
    mainPageUtils.isLoggedIn.value = value.data !== undefined
    isLoading.value = false;
  });
})

</script>