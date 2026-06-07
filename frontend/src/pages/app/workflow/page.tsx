import React from "react";
import { Box, Typography, Container, Paper } from "@mui/material";

const WorkflowPage: React.FC = () => {
  return (
    <Container maxWidth="xl">
      <Box sx={{ mt: 4, mb: 4 }}>
        <Paper sx={{ p: 4 }}>
          <Typography variant="h4" component="h1" gutterBottom>
            Workflow
          </Typography>
          <Typography variant="body1" color="text.secondary">
            NUSA Education Operating System - Workflow
          </Typography>
          <Typography variant="body2" sx={{ mt: 2 }}>
            This page is under construction. Coming soon.
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
};

export default WorkflowPage;
