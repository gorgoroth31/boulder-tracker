<template>Wir loggen dich gerade aus...</template>

<script setup lang="ts">
import {onMounted} from "vue";
import mainPageUtils from "@/utils/mainPageUtils";
import { useLogto } from "@logto/vue";

onMounted(() => {
    mainPageUtils.pageTitle.value = "Logout"

    const { signOut } = useLogto();
    signOut(`${import.meta.env.VITE_APP_URI}/about`)
})
</script>