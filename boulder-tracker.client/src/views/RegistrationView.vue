<template>
  <div class="text-h4">Registrierung</div>
  <v-divider thickness="3" class="my-3"></v-divider>
  <div class="text-h6">Du bist jetzt angemeldet! Wie dürfen wir dich nennen?</div>
  <v-text-field class="my-5" v-model="username"></v-text-field>
  <v-btn baseColor="green-darken-1" @click="submit">Speichern</v-btn>
  <v-snackbar v-model="snackbar">
    {{snackBarText}}
  </v-snackbar>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {createUserForClaim} from "@/api/user.api";
import {User} from "@/models/user";
import {useRouter} from "vue-router";

let username = ref('')
const router = useRouter()

const snackbar = ref(false)
const snackBarText = ref("Leider ist ein Fehler aufgetreten. Versuche es erneut")

async function submit() {
  // add check if username exists
  let newUser: User = {
    username: username.value
  }
  const result = await createUserForClaim(newUser)

  if (result.status === 201) {
    await router.push({path: "/"});
  } else if (result.status === 400) {
      snackbar.value = true
  }
}
</script>