<script setup lang="ts">
import type {Guild, Topic} from "@/store/types";
import AuthLayout from "@/layouts/AuthLayout.vue";
import LoadingLayout from "@/layouts/LoadingLayout.vue";
import {useGuildState, useUserState, useTopicsState} from "@/store";
import type {Ref} from "vue";
import axios from "axios";

const route = useRoute();

const userState = useUserState();
const guildState = useGuildState();
const topicsState = useTopicsState();

const modal: Ref<boolean> = ref(false);

const emoji: Ref<string> = ref("");
const name: Ref<string> = ref("");
const description: Ref<string> = ref("");

onMounted(async () => {
  const guildId = route.params.id as string;

  await guildState.fetchGuild(guildId);
  await topicsState.fetchTopics(guildId);
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
        <p class="active">Topics</p>
      </div>
      <div class="heading">
        <h1>Topics</h1>
        <button @click="modal = true">Create</button>
      </div>
      <div class="topics">
        <Topic
          v-for="topic in topicsState.topics as Topic[]"
          :emoji="topic.emoji"
          :name="topic.name"
          :description="topic.description"
        />
      </div>
    </div>
    <Transition>
      <div v-if="modal" class="modal-container" @click="modal = false">
        <div class="modal" @click.stop="">
          <h2>Create new topic</h2>
          <div class="field">
            <small>Emoji</small>
            <input v-model="emoji" type="text" placeholder="Topic emoji">
          </div>
          <div class="field">
            <small>Name</small>
            <input v-model="name" type="text" placeholder="Topic name">
          </div>
          <div class="field">
            <small>Description</small>
            <input v-model="description" type="text" placeholder="Topic description">
          </div>
          <div class="actions">
            <button class="inactive" @click="modal = false;">Cancel</button>
            <button class="active" @click="topicsState.createTopic(guildState.id, emoji, name, description); modal = false;">Create</button>
          </div>
        </div>
      </div>
    </Transition>
  </AuthLayout>
</template>

<style scoped lang="scss">
  .v-enter-active,
  .v-leave-active {
    transition: opacity 120ms ease-in-out;
  }

  .v-enter-from,
  .v-leave-to {
    opacity: 0;
  }

  .modal-container {
    width: 100%;
    height: 100vh;

    display: flex;
    align-items: center;
    justify-content: center;

    top: 0;
    position: fixed;

    background-color: var(--background-300-40);
    backdrop-filter: blur(24px);

    .modal {
      width: calc(552px + 24px * 2);

      display: flex;
      flex-direction: column;
      gap: 24px;

      background-color: var(--background-500);
      border-radius: 24px;
      // border: 2px solid var(--background-400);

      padding: 24px;

      h2 {
        color: var(--foreground-500);

        font: {
          size: 24px;
          weight: 500;
        };

        line-height: 48px;
      }

      small {
        color: var(--foreground-400);

        font: {
          size: 16px;
          weight: 500;
        };

        line-height: 24px;
      }

      .field {
        display: flex;
        flex-direction: column;
        gap: 12px;

        input {
          color: var(--foreground-300);

          font: {
            size: 16px;
            weight: 500;
          };

          line-height: 24px;

          background-color: var(--background-300);
          border: none;
          border-radius: 16px;

          outline: none;
          transition: 120ms ease-in-out;

          padding: 12px 24px;

          &::placeholder {
            color: var(--foreground-100);
          }

          &:focus {
            box-shadow: 0 0 0 2px var(--accent);
          }
        }
      }

      .actions {
        width: 100%;

        display: flex;
        justify-content: flex-end;
        gap: 12px;

        button {
          display: flex;
          gap: 12px;

          font: {
            size: 16px;
            weight: 600;
          };

          line-height: 24px;

          border: none;
          border-radius: 16px;

          cursor: pointer;
          transition: 120ms ease-in-out;

          padding: 12px 24px;

          &:hover {
            opacity: .8;
          }

          &.active {
            color: var(--background-500);

            background-color: var(--accent);
          }

          &.inactive {
            color: var(--foreground-100);

            background-color: var(--background-300);
          }
        }
      }
    }
  }

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

    .topics {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(264px, 1fr));
      grid-gap: 24px;
    }
  }
</style>