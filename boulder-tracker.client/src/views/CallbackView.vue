<template>
  <!-- When it's working in progress -->
  <p v-if="isLoading">Redirecting...</p>
</template>

<script setup lang="ts">
import { useHandleSignInCallback } from '@logto/vue';
import router from '@/router';

const { isLoading, isAuthenticated } = useHandleSignInCallback(() => {
  // Do something when finished, e.g. redirect to home page
  if (isAuthenticated) {
    router.push("/")
  } else {
    router.push("/about")
  }
});
</script>