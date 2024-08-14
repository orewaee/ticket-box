<script setup lang="ts">
import type {Guild, Topic} from "@/store/types";
import AuthLayout from "@/layouts/AuthLayout.vue";
import {useGuildState, useUserState, useTopicState, useTopicsState} from "@/store";

const route = useRoute();

const userState = useUserState();
const guildState = useGuildState();
const topicState = useTopicState();
const topicsState = useTopicsState();

onMounted(async () => {
  const guildId = route.params.id as string;
  const topicId = route.params["topic"] as string;

  await guildState.fetchGuild(guildId);
  await topicState.getTopicById(topicId);
});
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
        <NuxtLink :to="'/dashboard/' + guildState.id + '/topics'" class="inactive">
          <p>Topics</p>
        </NuxtLink>
        <p class="inactive">/</p>
        <p class="active">{{topicState.emoji}} {{topicState.name}}</p>
      </div>
      <div class="heading">
        <h1>Topic</h1>
        <button @click="topicsState.removeTopicById(topicState.id); topicState.reset(); navigateTo('/dashboard/' + guildState.id + '/topics');">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="var(--red-500-100)" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon icon-tabler icons-tabler-outline icon-tabler-trash">
            <path stroke="none" d="M0 0h24v24H0z" fill="none" />
            <path d="M4 7l16 0" />
            <path d="M10 11l0 6" />
            <path d="M14 11l0 6" />
            <path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" />
            <path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" />
          </svg>
          Delete
        </button>
      </div>
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

  .heading {
    display: flex;
    align-items: center;
    justify-content: space-between;

    h1 {
      color: var(--foreground-500);
    }

    button {
      display: flex;
      gap: 12px;

      color: var(--red-500-100);

      font: {
        size: 16px;
        weight: 600;
      };

      line-height: 24px;

      background-color: var(--red-500-20);
      border: none;
      border-radius: 16px;

      cursor: pointer;
      transition: 120ms ease-in-out;

      padding: 12px 24px;

      &:hover {
        background-color: var(--red-500-40);
      }
    }
  }
}
</style>