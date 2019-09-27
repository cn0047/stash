class Repo {
  static getOk() {
    return new Promise((resolve) => {
      setTimeout(resolve('Code: 200 MSG: OK.'), 1000);
    });
  }
}

module.exports = Repo;
