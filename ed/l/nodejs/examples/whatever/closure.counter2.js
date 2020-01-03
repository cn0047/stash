const c = {
  init: () => {
    this.c = 0;
  },
  inc: () => {
    this.c +=1;
    console.log(this.c);
  }
}

c.init();
c.inc();
c.inc();
c.inc();
c.init();
c.inc();
c.inc();
c.inc();
c.inc();
c.inc();
