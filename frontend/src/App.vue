<script lang="ts" setup>
import StoppedView from '@/components/views/StoppedView.vue'
import WorkingView from '@/components/views/WorkingView.vue'
import WorkEndView from '@/components/views/WorkEndView.vue'
import RestingView from '@/components/views/RestingView.vue'
import RestEndView from '@/components/views/RestEndView.vue'
import { useSession } from '@/composables/useSession'
import * as API from '@w/go/main/App'

;(window as any).api = API // expose go api methods for dev convenience

const { session } = useSession()

</script>

<template>
  <Transition name="view">
    <StoppedView v-if="session.state == 'Stopped'" />
    <WorkingView v-else-if="session.state == 'Working'" />
    <WorkEndView v-else-if="session.state == 'WorkEnd'" />
    <RestingView v-else-if="session.state == 'Resting'" />
    <RestEndView v-else-if="session.state == 'RestEnd'" />
  </Transition>
</template>
