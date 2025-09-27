<template>Wir loggen dich gerade aus...</template>

<script setup lang="ts">
import {onMounted} from "vue";
import {useAuth0} from "@auth0/auth0-vue";
import mainPageUtils from "@/utils/mainPageUtils";
const { logout } = useAuth0();


onMounted(() => {
    mainPageUtils.pageTitle.value = "Logout"
    logout({ logoutParams: { returnTo: window.location.origin } });
})
</script>