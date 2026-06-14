-- Create messages table for user messaging
CREATE TABLE messages (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  sender_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  receiver_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  subject VARCHAR(255),
  content TEXT NOT NULL,
  is_read BOOLEAN DEFAULT false,
  read_at TIMESTAMP,
  parent_message_id UUID REFERENCES messages(id),
  created_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP,
  CHECK (sender_id != receiver_id)
);

-- Indexes for messages table
CREATE INDEX idx_messages_sender_id ON messages(sender_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_messages_receiver_id ON messages(receiver_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_messages_is_read ON messages(is_read) WHERE deleted_at IS NULL;
CREATE INDEX idx_messages_parent_message_id ON messages(parent_message_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_messages_conversation ON messages(LEAST(sender_id, receiver_id), GREATEST(sender_id, receiver_id)) WHERE deleted_at IS NULL;
