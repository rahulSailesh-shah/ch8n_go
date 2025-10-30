import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectValue,
  SelectTrigger,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect } from "react";
import { useForm } from "react-hook-form";
import z from "zod";

const formSchema = z.object({
  endpoint: z.string().url({ message: "Please enter a valid URL" }),
  method: z.enum(["GET", "POST", "PUT", "DELETE", "PATCH"]),
  body: z.string().optional(),
  //   .refine(),
});

export type HTTPFormType = z.infer<typeof formSchema>;

interface HTTPRequestDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSubmit: (values: HTTPFormType) => void;
  defaultEndpoint?: string;
  defaultMethod?: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
  defaultBody?: string;
}

export const HTTPRequestDialog = ({
  open,
  onOpenChange,
  onSubmit,
  defaultEndpoint = "",
  defaultMethod = "GET",
  defaultBody = "",
}: HTTPRequestDialogProps) => {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      endpoint: defaultEndpoint,
      method: defaultMethod,
      body: defaultBody,
    },
  });

  useEffect(() => {
    if (open) {
      form.reset({
        endpoint: defaultEndpoint,
        method: defaultMethod,
        body: defaultBody,
      });
    }
  }, [open, defaultEndpoint, defaultMethod, defaultBody, form]);

  const watchMethod = form.watch("method");
  const showBodyField = ["POST", "PUT", "PATCH"].includes(watchMethod);

  const handleSubmit = (data: z.infer<typeof formSchema>) => {
    onSubmit(data);
    onOpenChange(false);
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>HTTP Request</DialogTitle>
          <DialogDescription>
            Configure settings for HTTP Request
          </DialogDescription>
        </DialogHeader>
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmit)}
            className="space-y-8 mt-4"
          >
            <FormField
              control={form.control}
              name="method"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Method</FormLabel>
                  <Select onValueChange={field.onChange} value={field.value}>
                    <FormControl>
                      <SelectTrigger className="w-full">
                        <SelectValue placeholder="Select method" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="GET">GET</SelectItem>
                      <SelectItem value="POST">POST</SelectItem>
                      <SelectItem value="PUT">PUT</SelectItem>
                      <SelectItem value="DELETE">DELETE</SelectItem>
                      <SelectItem value="PATCH">PATCH</SelectItem>
                    </SelectContent>
                  </Select>
                  <FormDescription className="text-xs text-muted-foreground">
                    The HTTP method to use for the request.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="endpoint"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Endpoint URL</FormLabel>
                  <FormControl>
                    <Input
                      placeholder="https://api.example.com/users/{{httpResponse.data.id}}"
                      {...field}
                    />
                  </FormControl>
                  <FormDescription className="text-xs text-muted-foreground">
                    Static URL or use {"{{vaiables}}"} for simple values or{" "}
                    {"{{json variable}}"} to stringify objects
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
            {showBodyField && (
              <FormField
                control={form.control}
                name="body"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Request Body</FormLabel>
                    <FormControl>
                      <Textarea
                        className="min-h-[120px] font-mono text-sm"
                        placeholder={
                          '{\n  "userId": "{{httpResponse.data.id}}",\n  "name": "{{user.name}}",\n  "items": {{httpResponse.data.items}}\n}'
                        }
                        {...field}
                      />
                    </FormControl>
                    <FormDescription className="text-xs text-muted-foreground">
                      JSON with template variables. Use {"{{variables}}"} for
                      simple values or {"{{json variable}}"} to stringify
                      objects
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            )}
            <DialogFooter className="mt-4">
              <Button type="submit">Save</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};
