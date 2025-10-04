<template>
  <div class="text-h5 gy-0">Wann warst du bouldern?</div>
</template>

<script setup lang="ts">
import {onMounted, Ref, ref} from "vue";
  import mainPageUtils from "./../utils/mainPageUtils";
  import {Session} from "node:inspector";
import {getCurrentInProgressSession} from "@/api/session.api";

  const session: Ref<Session | undefined> = ref();
  
  onMounted(() => {
    mainPageUtils.pageTitle.value = "Session hinzufügen";
    
    getCurrentInProgressSession().then(value => {
      session.value = value.data;
      console.log(session.value);
    })
  })
</script>