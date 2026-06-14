import React from 'react';
import {
  Box,
  Button,
  Card,
  CardContent,
  Container,
  TextField,
  Typography,
  Alert,
  CircularProgress,
  Grid,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Save as SaveIcon,
} from '@mui/icons-material';
import { useSchool } from '@/services/queries/SchoolsQueryService';
import { useUpdateSchool } from '@/services/commands/SchoolsCommandService';
import { useNavigate, useParams } from 'react-router-dom';
import { Formik, Form, Field, FormikHelpers } from 'formik';
import * as Yup from 'yup';

interface UpdateSchoolFormValues {
  name: string;
  code: string;
  address: string;
  city: string;
  state: string;
  country: string;
  postal_code: string;
  phone: string;
  email: string;
  website: string;
  is_active: boolean;
}

const validationSchema = Yup.object({
  name: Yup.string().required('School name is required'),
  code: Yup.string().required('School code is required'),
  email: Yup.string().email('Invalid email address'),
  website: Yup.string().url('Invalid website URL'),
});

const EditSchool = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [alert, setAlert] = React.useState<{ show: boolean; message: string; severity: 'success' | 'error' }>({
    show: false,
    message: '',
    severity: 'success',
  });

  const { data: school, isLoading, error } = useSchool(id!);

  const updateMutation = useUpdateSchool({
    onSuccess: () => {
      setAlert({ show: true, message: 'School updated successfully', severity: 'success' });
      setTimeout(() => navigate(`/dashboard/schools/${id}`), 1500);
    },
    onError: (error: any) => {
      setAlert({ show: true, message: error.message || 'Failed to update school', severity: 'error' });
    },
  });

  const handleSubmit = (values: UpdateSchoolFormValues, helpers: FormikHelpers<UpdateSchoolFormValues>) => {
    updateMutation.mutate({ id: id!, data: values });
    helpers.setSubmitting(false);
  };

  if (isLoading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="400px">
        <CircularProgress />
      </Box>
    );
  }

  if (error || !school) {
    return (
      <Container maxWidth="xl">
        <Alert severity="error" sx={{ mt: 3 }}>
          Failed to load school details
        </Alert>
      </Container>
    );
  }

  const initialValues: UpdateSchoolFormValues = {
    name: school.name,
    code: school.code,
    address: school.address || '',
    city: school.city || '',
    state: school.state || '',
    country: school.country || '',
    postal_code: school.postal_code || '',
    phone: school.phone || '',
    email: school.email || '',
    website: school.website || '',
    is_active: school.is_active,
  };

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate(`/dashboard/schools/${id}`)}
          sx={{ mb: 2 }}
        >
          Back to School Details
        </Button>
        <Typography variant="h4" component="h1" gutterBottom>
          Edit School
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Update the school information
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
          <Formik
            initialValues={initialValues}
            validationSchema={validationSchema}
            onSubmit={handleSubmit}
            enableReinitialize
          >
            {({ values, errors, touched, isSubmitting, handleChange, handleBlur }) => (
              <Form>
                <Grid container spacing={3}>
                  <Grid size={{ xs: 12, md: 6 }}>
                    <Field
                      as={TextField}
                      name="name"
                      label="School Name"
                      fullWidth
                      required
                      value={values.name}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.name && Boolean(errors.name)}
                      helperText={touched.name && errors.name}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12, md: 6 }}>
                    <Field
                      as={TextField}
                      name="code"
                      label="School Code"
                      fullWidth
                      required
                      value={values.code}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.code && Boolean(errors.code)}
                      helperText={touched.code && errors.code}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12 }}>
                    <Field
                      as={TextField}
                      name="address"
                      label="Address"
                      fullWidth
                      multiline
                      rows={2}
                      value={values.address}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.address && Boolean(errors.address)}
                      helperText={touched.address && errors.address}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12, sm: 6 }}>
                    <Field
                      as={TextField}
                      name="city"
                      label="City"
                      fullWidth
                      value={values.city}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.city && Boolean(errors.city)}
                      helperText={touched.city && errors.city}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12, sm: 6 }}>
                    <Field
                      as={TextField}
                      name="state"
                      label="State/Province"
                      fullWidth
                      value={values.state}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.state && Boolean(errors.state)}
                      helperText={touched.state && errors.state}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12, sm: 6 }}>
                    <Field
                      as={TextField}
                      name="country"
                      label="Country"
                      fullWidth
                      value={values.country}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.country && Boolean(errors.country)}
                      helperText={touched.country && errors.country}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12, sm: 6 }}>
                    <Field
                      as={TextField}
                      name="postal_code"
                      label="Postal Code"
                      fullWidth
                      value={values.postal_code}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.postal_code && Boolean(errors.postal_code)}
                      helperText={touched.postal_code && errors.postal_code}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12, sm: 6 }}>
                    <Field
                      as={TextField}
                      name="phone"
                      label="Phone"
                      fullWidth
                      value={values.phone}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.phone && Boolean(errors.phone)}
                      helperText={touched.phone && errors.phone}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12, sm: 6 }}>
                    <Field
                      as={TextField}
                      name="email"
                      label="Email"
                      fullWidth
                      type="email"
                      value={values.email}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.email && Boolean(errors.email)}
                      helperText={touched.email && errors.email}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12 }}>
                    <Field
                      as={TextField}
                      name="website"
                      label="Website"
                      fullWidth
                      value={values.website}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.website && Boolean(errors.website)}
                      helperText={touched.website && errors.website}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12 }}>
                    <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end', mt: 3 }}>
                      <Button
                        variant="outlined"
                        onClick={() => navigate(`/dashboard/schools/${id}`)}
                        disabled={isSubmitting}
                      >
                        Cancel
                      </Button>
                      <Button
                        type="submit"
                        variant="contained"
                        startIcon={isSubmitting ? <CircularProgress size={20} /> : <SaveIcon />}
                        disabled={isSubmitting}
                      >
                        {isSubmitting ? 'Saving...' : 'Save Changes'}
                      </Button>
                    </Box>
                  </Grid>
                </Grid>
              </Form>
            )}
          </Formik>
        </CardContent>
      </Card>
    </Container>
  );
};

export default EditSchool;
