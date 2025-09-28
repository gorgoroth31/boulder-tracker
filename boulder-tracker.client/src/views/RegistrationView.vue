<template>
  <div class="d-flex flex-column ga-5">
    <div class="text-h6">Du bist jetzt angemeldet! Wie dürfen wir dich nennen?</div>
    <v-text-field v-model="username" @input="handleInput()"></v-text-field>
    <div v-if="isError" class="text-warning">
      <v-icon>mdi-alert</v-icon> {{errorText}}
    </div>
    <v-btn baseColor="green-darken-1" @click="submit">Speichern</v-btn>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import {createUserForClaim} from "@/api/user.api";
import {User} from "@/models/user";
import {useRouter} from "vue-router";
import mainPageUtils from "@/utils/mainPageUtils";
import {existsUserNameAlready} from "@/api/api";

let username = ref('')
const router = useRouter()

const isError = ref(false)
const errorText = ref("Leider ist ein Fehler aufgetreten. Bitte versuche es erneut.")

onMounted(async () => {
  mainPageUtils.pageTitle.value = "Registrierung"
  mainPageUtils.isLoggedIn.value = false;
})

function handleInput() {
  errorText.value = "Leider ist ein Fehler aufgetreten. Bitte versuche es erneut.";
  isError.value = false;
}

async function submit() {
  if (username.value === "") {
    isError.value = true
    errorText.value = "Der Nutzername darf nicht leer sein."
    return;
  }
  
  let newUser: User = {
    username: username.value
  }
  const result = await createUserForClaim(newUser)

  if (result.status === 201) {
    await router.push({path: "/"});
    mainPageUtils.isLoggedIn.value = true;
  } else if (result.status === 400) {
    console.error(result)
    errorText.value = result.response.data;
    isError.value = true
  }
}
</script>