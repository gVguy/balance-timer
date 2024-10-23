<script lang="ts" setup>
import ViewLayout from '@/components/ViewLayout.vue'
import StopButton from '@/components/StopButton.vue'
import ContinueButton from '@/components/ContinueButton.vue'
import { useSession } from '@/composables/useSession'
import { computed } from 'vue'
import { BeginWork, Skip } from '@w/go/main/App'
import SkipButton from '@/components/SkipButton.vue'

const { session } = useSession()

const restTime = computed(() => {
  const { restMin } = session.value.config
  const s = restMin == 1 ? '' : 's'
  return `${restMin} minute${s}`
})

</script>

<template>
  <ViewLayout>
    <template #heading>Back to productivity</template>
    <ContinueButton @pointerdown="BeginWork" />
    <span>
      The time you've set aside for rest<br>has come to an end.
    </span>
    <span>
      Your next beak is scheduled in <span class="highlight">{{ restTime }}</span>.<br>
      And now, ease back into working on your tasks.
    </span>
    <template #actions>
      <SkipButton
        title="Skip productivity"
        text="Press and hold to move on to the next Rest period without going into a Productivity period"
        @confirm="Skip"
      />
      <StopButton />
    </template>
  </ViewLayout>
</template>

<style scoped lang="scss">
.countdown {
  color: rgba(var(--fg-rgb), .3);
}
</style>
