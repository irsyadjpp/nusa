/**
 * AI Assistant Component
 * AI-powered assistant for generating narrative content
 */

import { Box, Typography, Button, Stack, TextField, Chip } from '@mui/material';
import { AutoAwesome as AutoAwesomeIcon } from '@mui/icons-material';
import { useState } from 'react';

interface AIAssistantProps {
  onGenerate: (prompt: string) => void;
  loading?: boolean;
}

export const AIAssistant = ({ onGenerate, loading = false }: AIAssistantProps) => {
  const [prompt, setPrompt] = useState('');

  const suggestions = [
    'Generate a summary of student achievements',
    'Write about student strengths',
    'Describe areas for improvement',
    'Create a parent-friendly summary',
  ];

  return (
    <Box sx={{ p: 3, bgcolor: 'background.paper', borderRadius: 1, border: '1px solid divider' }}>
      <Stack direction="row" alignItems="center" gap={1} mb={2}>
        <AutoAwesomeIcon color="primary" />
        <Typography variant="h6">AI Assistant</Typography>
      </Stack>

      <Box sx={{ mb: 2 }}>
        <Typography variant="body2" color="textSecondary" gutterBottom>
          Quick suggestions:
        </Typography>
        <Stack direction="row" flexWrap="wrap" gap={1}>
          {suggestions.map((suggestion) => (
            <Chip
              key={suggestion}
              label={suggestion}
              size="small"
              onClick={() => setPrompt(suggestion)}
              clickable
            />
          ))}
        </Stack>
      </Box>

      <TextField
        fullWidth
        multiline
        rows={3}
        label="Describe what you want to generate"
        value={prompt}
        onChange={(e) => setPrompt(e.target.value)}
        placeholder="Enter your prompt..."
        sx={{ mb: 2 }}
      />

      <Button
        variant="contained"
        startIcon={<AutoAwesomeIcon />}
        onClick={() => onGenerate(prompt)}
        disabled={!prompt || loading}
        fullWidth
      >
        {loading ? 'Generating...' : 'Generate with AI'}
      </Button>
    </Box>
  );
};
