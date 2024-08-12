<script setup lang="ts">
defineProps<{
  id: string
  name: string
  iconUrl: string
  withBot: boolean
}>();
</script>

<template>
  <div class="guild">
    <img class="icon" draggable="false" :src="iconUrl" alt="icon">
    <div class="info">
      <h4>{{name}}</h4>
      <p v-if="withBot" class="ready">Ready</p>
      <p v-else class="need-a-bot">Need a bot</p>
    </div>
  </div>
</template>

<style scoped lang="scss">
  .guild {
    display: flex;
    gap: 12px;

    background-color: var(--background-300);
    border-radius: 24px;

    cursor: pointer;
    transition: 120ms ease-in-out;

    padding: 24px;

    &:hover {
      background-color: var(--background-200);
    }

    .icon {
      width: 48px;
      height: 48px;

      border-radius: 16px;

      user-select: none;
    }

    .info {
      display: flex;
      flex-direction: column;

      h4 {
        color: var(--foreground-400);

        font: {
          size: 16px;
          weight: 600;
        };

        line-height: 24px;
      }

      p {
        font: {
          size: 16px;
          weight: 600;
        };

        line-height: 24px;

        &.ready {
          color: var(--accent);
        }

        &.need-a-bot {
          color: var(--foreground-100);
        }
      }
    }
  }
</style>