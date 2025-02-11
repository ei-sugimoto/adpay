'use client';

import { Button } from '@/components/ui/button';
import { redirect } from 'next/navigation';
export default function NewProjectButton() {
  const redirectNewProject = () => {
    redirect('/dashboard/project/new');
  };
  return (
    <>
      <Button variant='outline' onClick={redirectNewProject}>
        New Project
      </Button>
    </>
  );
}
