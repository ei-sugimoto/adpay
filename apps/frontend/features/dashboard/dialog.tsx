import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTrigger,
} from '@/components/ui/dialog';
import { DialogTitle } from '@radix-ui/react-dialog';
import ProjectForm from './form';

export default function DashboardDialog() {
  return (
    <Dialog open>
      <DialogTrigger asChild>
        <Button variant='outline'>New Project</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>New Project</DialogTitle>
          <DialogDescription>
            Make a new project. This will create a new project in the database.
          </DialogDescription>
        </DialogHeader>
        <ProjectForm />
      </DialogContent>
    </Dialog>
  );
}
