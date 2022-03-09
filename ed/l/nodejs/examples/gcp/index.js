const {PubSub} = require('@google-cloud/pubsub');

const projectId = '';
const topicName = 'test-1';
const subscriptionName = 'test-1';

async function createTopic() {
  const pubSub = new PubSub({projectId});
  const [topic] = await pubSub.createTopic(topicName);

  console.log('created topic:', topic);
}

async function publish() {
  const pubSub = new PubSub({projectId});
  const topic = pubSub.topic(topicName);

  const message = {
    data: {
      message: "nodejs test",
    },
  };
  const messageBuffer = Buffer.from(JSON.stringify(message), 'utf8');

  try {
    await topic.publish(messageBuffer);
    console.log('Message published.');
  } catch (err) {
    console.error(err);
  }

}

// createTopic();
publish();
