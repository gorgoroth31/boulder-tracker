<template>
  <v-app :class="'overflow-scroll flex flex-col md:flex-row gap-4 ' + mainPageUtil.appCss.value">
      <div class="padding-1rem">
        <div class="text-h4 text-break">{{ mainPageUtil.pageTitle.value }}</div>
        <v-divider thickness="3" class="my-3"></v-divider>
        <RouterView/>
      </div>
      <v-bottom-navigation grow v-if="mainPageUtil.isLoggedIn.value">
        <v-btn to="/" value="home">
          <v-icon>mdi-home</v-icon>
          <span>Home</span>
        </v-btn>

        <v-btn to="/sessions/add" value="sessions/add">
          <v-icon>mdi-plus</v-icon>
          <span>Neue Session</span>
        </v-btn>
        
        <v-btn value="more">
          <v-icon>mdi-menu</v-icon>
          <span>More</span>
          <v-menu activator="parent" width="100vw">
            <v-list>
              <v-list-item to="/about">
                <v-list-item-title class="d-flex ga-2">
                  <v-icon>mdi-information-outline</v-icon><span>  Über</span>
                </v-list-item-title>
              </v-list-item>

              <v-list-item to="/logout">
                <v-list-item-title class="d-flex ga-2">
                  <v-icon>mdi-logout</v-icon><span>Logout</span>
                </v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </v-btn>
      </v-bottom-navigation>
  </v-app>
</template>

<script setup lang="ts">
import { RouterView} from 'vue-router'
import mainPageUtil from "./utils/mainPageUtils";
</script>