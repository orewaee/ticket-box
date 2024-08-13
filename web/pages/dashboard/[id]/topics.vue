<script setup lang="ts">
import type {Guild} from "@/store/types";
import AuthLayout from "@/layouts/AuthLayout.vue";
import LoadingLayout from "@/layouts/LoadingLayout.vue";
import {useGuildState, useUserState} from "@/store";

const route = useRoute();

const userState = useUserState();
const guildState = useGuildState();

onMounted(async () => {
  await guildState.fetchGuild(route.params.id as string);
})
</script>

<template>
  <AuthLayout>
    <MyHeader :guild="guildState.guild as Guild" :avatar-url="userState.avatarUrl" />
    <div class="container">
      <div class="breadcrumbs">
        <NuxtLink to="/dashboard" class="inactive">
          <p>Dashboard</p>
        </NuxtLink>
        <p class="inactive">/</p>
        <NuxtLink :to="'/dashboard/' + guildState.id" class="inactive">
          <p>{{guildState.name}}</p>
        </NuxtLink>
        <p class="inactive">/</p>
        <p class="active">Topics</p>
      </div>
      <h1>Topics</h1>
    </div>
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
  }
</style>