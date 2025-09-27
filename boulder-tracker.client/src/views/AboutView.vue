<template>
  <div class="text-h5">Hier kommt eine putzig kleine About-page hin.</div>
  <Login v-if="!isLoggedIn"></Login>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import mainPageUtils from "@/utils/mainPageUtils";
import Login from "@/components/auth/Login.vue";
import {isUserAuthenticated} from "@/plugins/auth";

const isLoggedIn = ref<boolean>(false);

onMounted(async () => {
  mainPageUtils.pageTitle.value = "Über"
  isLoggedIn.value = await isUserAuthenticated()

})
</script>