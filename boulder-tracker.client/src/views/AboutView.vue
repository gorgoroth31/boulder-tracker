<template>
  <div class="d-flex flex-column gap-2rem">
    <v-alert color="amber">Diese App ist aktuell noch in der Beta-Version. Beim Release zu v1.0 werden möglicherweise
      alle Daten gelöscht. Näheres dazu im Github
    </v-alert>
    <div>
      <div class="text-h5 text-primary">Der Bouldertracker:</div>
      <div class="text-body-1">das neue Schweizer Taschenmesser, um deine Bouldersessions auf ein neues Level zu heben!</div>
    </div>

    <AboutCarousel :slides="slides1"/>

    <v-btn v-if="!mainPageUtils.isLoggedIn.value" @click="login">Jetzt loslegen</v-btn>
    <div>
      <div class="text-h5 text-primary">Schwierigkeitsskalen:</div>
      <div class="text-body-1">Aktuell wird nur die Schwierigkeitsskala der
        <a class="text-secondary link-underline-none" href="https://bloc-huette.de/">Blochütte Augsburg</a> unterstützt.
        Weitere Skalen werden im Lauf der Zeit hinzugefügt und wer weiß, vielleicht könnt ihr
        ja eines Tages selbst die Skalen eures Gyms einreichen ;)
      </div>
    </div>

    <AboutCarousel :slides="slides2"/>

    <div>
      <div class="text-h5 text-primary">Open Source:</div>
      <div class="text-body1">Ja genau, der Bouldertracker ist komplett Open Source und als Hobbyprojekt von mir
        entstanden. Falls du näheres über mich oder das Projekt erfahren willst, findest du unten den Github Link.<br/>
        Dort kannst du übrigens auch Features anfragen oder Bugs einreichen, die du während der Nutzung gefunden hast.
      </div>
    </div>
    
    <v-btn class="w-50" href="https://github.com/gorgoroth31/boulder-tracker"><v-icon>mdi-github</v-icon><div class="ml-2">Github</div></v-btn>
  </div>
</template>

<script setup lang="ts">
import {onMounted} from "vue";
import mainPageUtils from "@/utils/mainPageUtils";
import {loginWithRedirect} from "@/plugins/auth";
import AboutCarousel from "@/components/about/AboutCarousel.vue";

// TODO: anstatt dem Text schicke Grafiken basteln, die die App als Showcase wie z.B. im Play Store anzeigen
const slides1 = [
  'Sessions anlegen',
  'Statistiken betrachten',
]
const slides2 = [
  'Open Source und ohne Werbung',
  'Features und Bugs selbst einreichen'
]

onMounted(async () => {
  mainPageUtils.pageTitle.value = "Über die App";
})

function login() {
  loginWithRedirect()
}
</script>