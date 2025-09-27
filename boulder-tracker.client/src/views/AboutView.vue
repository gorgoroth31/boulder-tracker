<template class="h-100 d-flex flex-column justify-content-between">
  <div class="text-h5">
    Hier kommt eine putzig kleine About-page hin.<br/><br/>
    Aktuell hat die App noch kaum Funktionalität, das wird sich aber mit der Zeit ändern. Schau einfach öfter mal vorbei
    :D
  </div>
  <v-btn @click="stuff">Stuff</v-btn>
  <div class="flex-grow-1"></div>
  <v-btn v-if="!mainPageUtils.isLoggedIn.value" @click="login">Anmelden</v-btn>
</template>

<script setup lang="ts">
import {onMounted} from "vue";
import mainPageUtils from "@/utils/mainPageUtils";
import {loginWithRedirect} from "@/plugins/auth";

onMounted(async () => {
  mainPageUtils.pageTitle.value = "Über"
})

function login() {
  loginWithRedirect()
}

function stuff() {
  console.log(import.meta.env)
}
</script>