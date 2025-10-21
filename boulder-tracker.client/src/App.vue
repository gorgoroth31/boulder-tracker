<template>
  <v-app :class="'overflow-scroll flex flex-col md:flex-row gap-4 ' + mainPageUtil.appCss.value">
    <div class="w-100 bg-amber text-center">Beta-Version</div>
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
</script>