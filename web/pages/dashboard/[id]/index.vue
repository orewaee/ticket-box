<script setup lang="ts">
import {useUserState, useGuildState} from "@/store";
import type {Guild} from "@/store/types";
import AuthLayout from "@/layouts/AuthLayout.vue";
import LoadingLayout from "@/layouts/LoadingLayout.vue";

const route = useRoute();

const userState = useUserState();
const guildState = useGuildState();

onMounted(async () => {
  await guildState.fetchGuild(route.params.id as string);
});
</script>

<template>
  <AuthLayout>
    <LoadingLayout :loading="guildState.loading">
      <MyHeader :guild="guildState.guild as Guild" :avatar-url="userState.avatarUrl" />
      <div class="container">
        <div class="breadcrumbs">
          <NuxtLink to="/dashboard" class="inactive">
            <p>Dashboard</p>
          </NuxtLink>
          <p class="inactive">/</p>
          <p class="active">{{guildState.name}}</p>
        </div>
        <h1>Config</h1>
        <a href="https://discord.com/oauth2/authorize?client_id=1207683042680897629&permissions=8&integration_type=0&scope=bot+applications.commands" v-if="!(guildState.guild as Guild).with_bot">
          <ConfigOption title="Add a bot to the server" description="There is no bot on this server, invite a bot to get started">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="var(--foreground-100)" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon icon-tabler icons-tabler-outline icon-tabler-arrow-up-right">
              <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
              <path d="M17 7l-10 10" />
              <path d="M8 7l9 0l0 9" />
            </svg>
          </ConfigOption>
        </a>

        <NuxtLink :to="'/dashboard/' + guildState.id + '/topics'" v-if="(guildState.guild as Guild).with_bot">
          <ConfigOption title="Topics" description="List of guild topics">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="var(--foreground-100)" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon icon-tabler icons-tabler-outline icon-tabler-arrow-right">
              <path stroke="none" d="M0 0h24v24H0z" fill="none" />
              <path d="M5 12l14 0" />
              <path d="M13 18l6 -6" />
              <path d="M13 6l6 6" />
            </svg>
          </ConfigOption>
        </NuxtLink>
      </div>
    </LoadingLayout>
  </AuthLayout>
</template>

<style scoped lang="scss">
  .container {
    max-width: calc(840px + 24px * 2);

    display: flex;
    flex-direction: column;
    gap: 24px;

    margin: 0 auto;
    padding: 24px;

    .breadcrumbs {
      display: flex;
      gap: 12px;

      font: {
        size: 16px;
        weight: 500;
      };

      line-height: 24px;

      .inactive {
        color: var(--foreground-200);

        text-decoration: none;
      }

      .active {
        color: var(--foreground-400);
      }
    }

    h1 {
      color: var(--foreground-500);
    }

    a {
      text-decoration: none;
    }
  }
</style>