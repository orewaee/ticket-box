<script setup lang="ts">
import {useUserState, useGuildsState} from "@/store";
import AuthLayout from "@/layouts/AuthLayout.vue";
import Guild from "@/components/Guild.vue";

const userState = useUserState();
const guildsState = useGuildsState();

onMounted(async () => {
  await guildsState.fetchGuilds();
});
</script>

<template>
  <AuthLayout>
    <MyHeader :guild="null" :avatar-url="userState.avatarUrl" />
    <div class="container">
      <h1>Dashboard</h1>
      <div class="guilds-skeleton" v-if="guildsState.loading">
        <div class="item" />
        <div class="item" />
        <div class="item" />
      </div>

      <div class="guilds" v-else>
        <NuxtLink class="link" v-for="guild in guildsState.guilds as Guild[]" :to="'/dashboard/' + guild.id">
          <Guild :id="guild.id" :with-bot="guild.with_bot" :name="guild.name" :icon-url="guild.icon_url" />
        </NuxtLink>
      </div>
    </div>
  </AuthLayout>
</template>

<style scoped lang="scss">
  p {
    color: var(--foreground-400);
  }

  .container {
    max-width: calc(840px + 24px * 2);

    display: flex;
    flex-direction: column;
    gap: 24px;

    margin: 0 auto;
    padding: 24px;

    .guilds-skeleton {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(408px, 1fr));
      grid-gap: 24px;

      .item {
        height: 96px;

        background: linear-gradient(
          135deg,
          var(--background-300),
          var(--background-400),
          var(--background-300)
        );

        background-size: 200% 100%;
        border-radius: 24px;

        animation: 2s shine linear infinite;
      }

      @keyframes shine {
        to {
          background-position-x: -200%;
        }
      }
    }

    h1 {
      color: var(--foreground-500);
    }

    .guilds {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(408px, 1fr));
      grid-gap: 24px;

      .link {
        text-decoration: none;
      }
    }
  }
</style>