import React, { useState } from 'react';
import {
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  Container,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  TablePagination,
  TextField,
  Typography,
  Menu,
  MenuItem,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Alert,
  CircularProgress,
} from '@mui/material';
import {
  MoreVert as MoreVertIcon,
  Add as AddIcon,
  Edit as EditIcon,
  Delete as DeleteIcon,
  Search as SearchIcon,
  Visibility as VisibilityIcon,
} from '@mui/icons-material';
import { useSchools } from '@/services/queries/SchoolsQueryService';
import { useDeleteSchool, useUpdateSchoolStatus } from '@/services/commands/SchoolsCommandService';
import { useNavigate } from 'react-router-dom';

const SchoolsList = () => {
  const navigate = useNavigate();
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState('');
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [selectedSchool, setSelectedSchool] = useState<any>(null);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [alert, setAlert] = useState<{ show: boolean; message: string; severity: 'success' | 'error' }>({
    show: false,
    message: '',
    severity: 'success',
  });

  const { data: schoolsData, isLoading, error, refetch } = useSchools({
    page: page + 1,
    page_size: rowsPerPage,
    search: search || undefined,
    is_active: statusFilter ? statusFilter === 'active' : undefined,
  });

  const deleteMutation = useDeleteSchool({
    onSuccess: () => {
      setAlert({ show: true, message: 'School deleted successfully', severity: 'success' });
      refetch();
      setDeleteDialogOpen(false);
    },
    onError: (error: any) => {
      setAlert({ show: true, message: error.message || 'Failed to delete school', severity: 'error' });
    },
  });

  const statusMutation = useUpdateSchoolStatus({
    onSuccess: () => {
      setAlert({ show: true, message: 'School status updated successfully', severity: 'success' });
      refetch();
    },
    onError: (error: any) => {
      setAlert({ show: true, message: error.message || 'Failed to update school status', severity: 'error' });
    },
  });

  const handleMenuClick = (event: React.MouseEvent<HTMLElement>, school: any) => {
    setAnchorEl(event.currentTarget);
    setSelectedSchool(school);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
    setSelectedSchool(null);
  };

  const handleView = () => {
    navigate(`/dashboard/schools/${selectedSchool?.id}`);
    handleMenuClose();
  };

  const handleEdit = () => {
    navigate(`/dashboard/schools/${selectedSchool?.id}`);
    handleMenuClose();
  };

  const handleDeleteClick = () => {
    setDeleteDialogOpen(true);
    handleMenuClose();
  };

  const handleDeleteConfirm = () => {
    if (selectedSchool) {
      deleteMutation.mutate(selectedSchool.id);
    }
  };

  const handleToggleStatus = () => {
    if (selectedSchool) {
      statusMutation.mutate({ id: selectedSchool.id, is_active: !selectedSchool.is_active });
    }
    handleMenuClose();
  };

  const handleChangePage = (_event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  const handleSearch = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(event.target.value);
    setPage(0);
  };

  if (isLoading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="400px">
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Container maxWidth="xl">
        <Alert severity="error" sx={{ mt: 3 }}>
          Failed to load schools: {(error as Error).message}
        </Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Schools Management
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Manage all schools in the system
        </Typography>
      </Box>

      {alert.show && (
        <Alert
          severity={alert.severity}
          onClose={() => setAlert({ ...alert, show: false })}
          sx={{ mb: 2 }}
        >
          {alert.message}
        </Alert>
      )}

      <Card>
        <CardContent>
          <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
            <Box sx={{ display: 'flex', gap: 2, alignItems: 'center' }}>
              <TextField
                placeholder="Search schools..."
                size="small"
                value={search}
                onChange={handleSearch}
                InputProps={{
                  startAdornment: <SearchIcon sx={{ mr: 1, color: 'text.secondary' }} />,
                }}
                sx={{ width: 300 }}
              />
              <TextField
                select
                size="small"
                value={statusFilter}
                onChange={(e) => setStatusFilter(e.target.value)}
                sx={{ width: 150 }}
              >
                <MenuItem value="">All Status</MenuItem>
                <MenuItem value="active">Active</MenuItem>
                <MenuItem value="inactive">Inactive</MenuItem>
              </TextField>
            </Box>
            <Button
              variant="contained"
              startIcon={<AddIcon />}
              onClick={() => navigate('/dashboard/schools/new')}
            >
              Add School
            </Button>
          </Box>

          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Name</TableCell>
                  <TableCell>Code</TableCell>
                  <TableCell>City</TableCell>
                  <TableCell>Phone</TableCell>
                  <TableCell>Email</TableCell>
                  <TableCell>Status</TableCell>
                  <TableCell>Created At</TableCell>
                  <TableCell align="right">Actions</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {schoolsData?.schools?.map((school: any) => (
                  <TableRow key={school.id} hover>
                    <TableCell>{school.name}</TableCell>
                    <TableCell>{school.code}</TableCell>
                    <TableCell>{school.city || '-'}</TableCell>
                    <TableCell>{school.phone || '-'}</TableCell>
                    <TableCell>{school.email || '-'}</TableCell>
                    <TableCell>
                      <Chip
                        label={school.is_active ? 'Active' : 'Inactive'}
                        color={school.is_active ? 'success' : 'default'}
                        size="small"
                      />
                    </TableCell>
                    <TableCell>
                      {new Date(school.created_at).toLocaleDateString()}
                    </TableCell>
                    <TableCell align="right">
                      <IconButton
                        size="small"
                        onClick={(e) => handleMenuClick(e, school)}
                      >
                        <MoreVertIcon />
                      </IconButton>
                    </TableCell>
                  </TableRow>
                ))}
                {schoolsData?.schools?.length === 0 && (
                  <TableRow>
                    <TableCell colSpan={8} align="center" sx={{ py: 3 }}>
                      <Typography color="text.secondary">No schools found</Typography>
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </TableContainer>

          <TablePagination
            component="div"
            count={schoolsData?.total || 0}
            page={page}
            onPageChange={handleChangePage}
            rowsPerPage={rowsPerPage}
            onRowsPerPageChange={handleChangeRowsPerPage}
            rowsPerPageOptions={[5, 10, 25, 50]}
          />
        </CardContent>
      </Card>

      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleMenuClose}
      >
        <MenuItem onClick={handleView}>
          <VisibilityIcon sx={{ mr: 1 }} fontSize="small" />
          View Details
        </MenuItem>
        <MenuItem onClick={handleEdit}>
          <EditIcon sx={{ mr: 1 }} fontSize="small" />
          Edit
        </MenuItem>
        <MenuItem onClick={handleToggleStatus}>
          {selectedSchool?.is_active ? 'Deactivate' : 'Activate'}
        </MenuItem>
        <MenuItem onClick={handleDeleteClick} sx={{ color: 'error.main' }}>
          <DeleteIcon sx={{ mr: 1 }} fontSize="small" />
          Delete
        </MenuItem>
      </Menu>

      <Dialog open={deleteDialogOpen} onClose={() => setDeleteDialogOpen(false)}>
        <DialogTitle>Confirm Delete</DialogTitle>
        <DialogContent>
          <Typography>
            Are you sure you want to delete the school "{selectedSchool?.name}"? This action cannot be undone.
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setDeleteDialogOpen(false)}>Cancel</Button>
          <Button
            onClick={handleDeleteConfirm}
            color="error"
            disabled={deleteMutation.isPending}
          >
            {deleteMutation.isPending ? <CircularProgress size={20} /> : 'Delete'}
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default SchoolsList;
