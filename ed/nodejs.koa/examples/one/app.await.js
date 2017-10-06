const Koa = require('koa');
const Router = require('koa-router');

const Repo = require('./repo');

const app = new Koa();
const router = new Router({
  prefix: '/v1'
});

router.get('/getOk', async (ctx, next) => {
  ctx.body = await Repo.getOk();
});

app.use(router.routes());
app.listen(3000);
