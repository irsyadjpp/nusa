import React, { useState } from 'react';
import { Box, Grid, Paper, Button, Typography, Container, Breadcrumbs, Link, Divider } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { useATPSets } from '@/services/queries/ATPQueryService';
import { ATPSet } from '@/shared/types/domain';

const ATPSetPage: React.FC = () => {
  const navigate = useNavigate();
  const [selectedATPSet, setSelectedATPSet] = useState<ATPSet | null>(null);
  const [filters] = useState<{ tp_set_id?: string; status?: string }>({});

  // Query ATP Sets
  const { data: atpSets = [], isLoading } = useATPSets(filters);

  const handleCreateNew = () => {
    setSelectedATPSet(null);
    navigate('/atp-sets/create');
  };

  const handleSelectATPSet = (atpSet: ATPSet) => {
    setSelectedATPSet(atpSet);
  };

  const handleEditATPSet = () => {
    if (selectedATPSet) {
      navigate(`/atp-sets/${selectedATPSet.id}/edit`);
    }
  };

  const handleDeleteATPSet = () => {
    // TODO: Implement delete functionality
    if (selectedATPSet && window.confirm('Are you sure you want to delete this ATP Set?')) {
      console.log('Delete ATP Set:', selectedATPSet.id);
    }
  };

  const handleApproveATPSet = () => {
    // TODO: Implement approve functionality
    if (selectedATPSet && window.confirm('Are you sure you want to approve this ATP Set?')) {
      console.log('Approve ATP Set:', selectedATPSet.id);
    }
  };

  const handleCreateATPFromSet = () => {
    if (selectedATPSet) {
      navigate('/atp/create', { state: { selectedATPSet } });
    }
  };

  // Simple list of ATP sets (reusing ATPList component structure)
  const renderATPSetList = () => {
    if (isLoading) {
      return <Typography>Loading ATP Sets...</Typography>;
    }

    if (!atpSets || atpSets.length === 0) {
      return <Typography>No ATP Sets found</Typography>;
    }

    return (
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        {atpSets.map((atpSet) => (
          <Paper
            key={atpSet.id}
            sx={{
              p: 2,
              cursor: 'pointer',
              border: selectedATPSet?.id === atpSet.id ? 2 : 1,
              borderColor: selectedATPSet?.id === atpSet.id ? 'primary.main' : 'divider',
            }}
            onClick={() => handleSelectATPSet(atpSet as any)}
          >
            <Typography variant="h6">{atpSet.id}</Typography>
            <Typography variant="body2" color="text.secondary">
              TP Set: {atpSet.tp_set_id}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              Status: {atpSet.status}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              Created: {new Date(atpSet.created_at).toLocaleDateString()}
            </Typography>
          </Paper>
        ))}
      </Box>
    );
  };

  return (
    <Container maxWidth="xl">
      <Box sx={{ py: 3 }}>
        <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
          <Link underline="hover" color="inherit" onClick={() => navigate('/learning-design')}>
            Learning Design
          </Link>
          <Typography color="text.primary">ATP Sets</Typography>
        </Breadcrumbs>

        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
          <Typography variant="h4" component="h1">
            ATP Sets
          </Typography>
          <Button
            variant="contained"
            onClick={handleCreateNew}
          >
            Create ATP Set
          </Button>
        </Box>

        <Grid container spacing={3}>
          {/* Left Panel - ATP Set List */}
          <Grid size={{ xs: 12, md: selectedATPSet ? 6 : 12 }}>
            <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
              <Box sx={{ p: 2 }}>
                <Typography variant="h6" gutterBottom>
                  Filters
                </Typography>
                {/* Simple TP Set filter */}
                <Box sx={{ mt: 2 }}>
                  <Typography variant="body2" color="text.secondary" gutterBottom>
                    Filter by TP Set or status coming soon
                  </Typography>
                </Box>
              </Box>
              <Divider />
              <Box sx={{ flex: 1, overflow: 'auto', p: 2 }}>
                {renderATPSetList()}
              </Box>
            </Paper>
          </Grid>

          {/* Right Panel - ATP Set Detail */}
          {selectedATPSet && (
            <Grid size={{ xs: 12, md: 6 }}>
              <Paper sx={{ height: '100%', p: 2 }}>
                <Typography variant="h6" gutterBottom>
                  ATP Set Detail
                </Typography>
                <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      ID
                    </Typography>
                    <Typography variant="body1">
                      {selectedATPSet.id}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      TP Set ID
                    </Typography>
                    <Typography variant="body1">
                      {selectedATPSet.tp_set_id}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Subject
                    </Typography>
                    <Typography variant="body1">
                      {selectedATPSet.subject_id}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Phase
                    </Typography>
                    <Typography variant="body1">
                      {selectedATPSet.phase_id}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Grade
                    </Typography>
                    <Typography variant="body1">
                      {selectedATPSet.grade}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Semester
                    </Typography>
                    <Typography variant="body1">
                      {selectedATPSet.semester}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Status
                    </Typography>
                    <Typography variant="body1">
                      {selectedATPSet.status}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Created By
                    </Typography>
                    <Typography variant="body1">
                      {selectedATPSet.created_by}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Created
                    </Typography>
                    <Typography variant="body1">
                      {new Date(selectedATPSet.created_at).toLocaleString()}
                    </Typography>
                  </Box>
                  {selectedATPSet.approved_at && (
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        Approved
                      </Typography>
                      <Typography variant="body1">
                        {new Date(selectedATPSet.approved_at).toLocaleString()}
                      </Typography>
                    </Box>
                  )}
                </Box>
                
                <Divider sx={{ my: 2 }} />
                
                <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                  <Button
                    variant="contained"
                    onClick={handleCreateATPFromSet}
                  >
                    Create ATP from this Set
                  </Button>
                  <Button
                    variant="outlined"
                    onClick={handleEditATPSet}
                  >
                    Edit ATP Set
                  </Button>
                  {selectedATPSet.status !== 'APPROVED' && (
                    <Button
                      variant="outlined"
                      color="success"
                      onClick={handleApproveATPSet}
                    >
                      Approve ATP Set
                    </Button>
                  )}
                  <Button
                    variant="outlined"
                    color="error"
                    onClick={handleDeleteATPSet}
                  >
                    Delete ATP Set
                  </Button>
                </Box>
              </Paper>
            </Grid>
          )}
        </Grid>
      </Box>
    </Container>
  );
};

export default ATPSetPage;