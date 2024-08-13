<script setup lang="ts">
defineProps<{
  emoji: string
  name: string
  description: string
}>();
</script>

<template>
  <div class="topic">
    <div class="name">
      <h4>{{emoji}}</h4>
      <h4>{{name}}</h4>
    </div>
    <p>{{description}}</p>
  </div>
</template>

<style scoped lang="scss">
  .topic {
    display: flex;
    flex-direction: column;
    gap: 12px;

    font: {
      size: 16px;
      weight: 600;
    };

    line-height: 24px;

    background-color: var(--background-300);
    border-radius: 24px;

    padding: 24px;

    .name {
      display: flex;
      flex-direction: row;
      gap: 12px;

      h4 {
        color: var(--foreground-400);
      }
    }



    p {
      color: var(--foreground-200);
    }
  }
</style>