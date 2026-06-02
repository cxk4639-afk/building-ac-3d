<script setup lang="ts">
defineProps<{
  embedded?: boolean
}>()
</script>

<template>
  <n-card title="3D配置">
    <n-space vertical size="large">
      <n-alert type="info" title="工程骨架">
        后续承载楼宇模型、楼层、房间、空调设备、点位绑定和配置保存。
      </n-alert>
      <n-data-table :columns="[]" :data="[]" />
    </n-space>
  </n-card>
</template>
