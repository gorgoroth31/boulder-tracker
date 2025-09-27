<template>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import {User} from "@/models/user";
import {getCurrentLoggedInUser} from "@/api/user.api";
import mainPageUtils from "./../utils/mainPageUtils";

const user = ref<User | null>(null); 

onMounted(async () => {
  await getCurrentLoggedInUser().then(value => {
    user.value = value.data;
    
    mainPageUtils.pageTitle.value = "Hallo " + user.value.userName
  });
})

</script>