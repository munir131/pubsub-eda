#!/usr/bin/env node

const { PubSub } = require('@google-cloud/pubsub');
const yargs = require('yargs');

// Replace with your topic and subscription names
const topicName = 'test';
const subscriptionName = 'test-sub';

// Initialize Pub/Sub client
const pubSubClient = new PubSub();

/**
 * Publishes a message to the Pub/Sub topic.
 * @param {string} message The message to publish
 */
async function publishMessage(message) {
  try {
    const messageId = await pubSubClient.topic(topicName).publishMessage({
      data: Buffer.from(message),
    });
    console.log(`Message published with ID: ${messageId}`);
  } catch (error) {
    console.error(`Error publishing message: ${error.message}`);
  }
}

/**
 * Listens for messages on the Pub/Sub subscription.
 */
function listenForMessages() {
  const subscription = pubSubClient.subscription(subscriptionName);

  subscription.on('message', (message) => {
    console.log(`Received message:`);
    console.log(`\tID: ${message.id}`);

    try {
      // Attempt to decode the message data
      const decodedData = Buffer.from(message.data, 'base64').toString('utf-8');
      // console.log(`\tData (decoded): ${decodedData}`);
      const parsedData = JSON.parse(decodedData);
      console.log("Decoded and parsed JSON object:", parsedData);
    } catch (error) {
      console.log(`\tRaw Data: ${message.data}`);
    }

    console.log(`\tAttributes: ${JSON.stringify(message.attributes)}`);
    message.ack(); // Acknowledge the message
  });

  subscription.on('error', (error) => {
    console.error(`Error listening for messages: ${error.message}`);
  });

  console.log(`Listening for messages on subscription "${subscriptionName}"...`);
}

// Define the CLI commands using yargs
yargs
  .command(
    'publish <message>',
    'Publish a message to the Pub/Sub topic',
    (yargs) => {
      yargs.positional('message', {
        describe: 'The message to publish',
        type: 'string',
      });
    },
    (argv) => {
      publishMessage(argv.message);
    }
  )
  .command(
    'listen',
    'Listen for messages from the Pub/Sub subscription',
    () => {},
    () => {
      listenForMessages();
    }
  )
  .demandCommand(1, 'You must provide a valid command.')
  .help()
  .wrap(null)
  .argv;
