<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Capitalism is just a game of luck and math! Get good.</h1>

        <form v-on:submit.prevent="findCompany">
          <div class="form-group">
            <input
              v-model="websiteUrl"
              type="text"
              id="website-input"
              placeholder="Enter a publicly-traded company name"
              class="form-control"
            />
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Find!</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
      websiteUrl: "",
    };
  },

  methods: {
    findCompany() {
      console.log(
        `I should find data so this can actually find ${this.websiteUrl}`
      );
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
