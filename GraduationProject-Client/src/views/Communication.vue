<template>
  <div class="communication">
    <div>
      <va-chat-box></va-chat-box>
    </div>
  </div>
</template>
<script>
import VAChatBox from "../widgets/VAChatBox.vue";
export default {
  name: "VACommunication",
  data() {
    return {
    };
  },
  components: {
    "va-chat-box": VAChatBox
  },
};
</script>