<script setup lang="ts">
import type {Guild} from "@/store/types";
import {useTokenState, useUserState} from "@/store";

const tokenState = useTokenState();
const userState = useUserState();

const props = defineProps<{
  guild: Guild | null
  avatarUrl: string
}>()
</script>

<template>
  <header>
    <div class="container">
      <div v-if="props.guild != null" class="guild">
        <img :src="guild?.icon_url" alt="icon" :draggable="false">
        <button>
          {{guild?.name}}
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="var(--foreground-100)" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon icon-tabler icons-tabler-outline icon-tabler-selector">
            <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
            <path d="M8 9l4 -4l4 4" />
            <path d="M16 15l-4 4l-4 -4" />
          </svg>
        </button>
      </div>
      <div v-else class="void" />
      <div class="user">
        <button @click="tokenState.removeToken(); userState.removeUser();">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="var(--foreground-100)" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon icon-tabler icons-tabler-outline icon-tabler-logout-2">
            <path stroke="none" d="M0 0h24v24H0z" fill="none" />
            <path d="M10 8v-2a2 2 0 0 1 2 -2h7a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-7a2 2 0 0 1 -2 -2v-2" />
            <path d="M15 12h-12l3 -3" /><path d="M6 15l-3 -3" />
          </svg>
        </button>
        <img :src="props.avatarUrl" alt="avatar" :draggable="false">
      </div>
    </div>
  </header>
</template>

<style scoped lang="scss">
  header {
    background-color: var(--background-500-40);
    backdrop-filter: blur(24px);
    border-bottom: 2px solid var(--background-300-40);

    position: sticky;
    top: 0;

    .container {
      max-width: calc(840px + 24px * 2);
      height: 96px;

      display: flex;
      justify-content: space-between;

      margin: 0 auto;
      padding: 24px;

      .guild {
        display: flex;
        gap: 12px;

        user-select: none;

        img {
          width: 48px;
          height: 48px;

          border-radius: 16px;
        }

        button {
          display: flex;
          gap: 12px;

          color: var(--foreground-300);

          font: {
            size: 16px;
            weight: 600;
          };

          line-height: 24px;

          background-color: var(--background-300);
          border: none;
          border-radius: 16px;

          cursor: pointer;
          transition: 120ms ease-in-out;

          padding: 12px 24px;

          &:hover {
            background-color: var(--background-200);
          }
        }
      }
    }

    .user {
      display: flex;
      gap: 12px;

      button {
        width: 48px;
        height: 48px;

        display: flex;
        align-items: center;
        justify-content: center;

        background-color: var(--background-300);
        border: none;
        border-radius: 16px;

        cursor: pointer;
        transition: 120ms ease-in-out;

        &:hover {
          background-color: var(--background-200);
        }
      }

      img {
        width: 48px;
        height: 48px;

        border-radius: 16px;

        user-select: none;
      }
    }
  }
</style>